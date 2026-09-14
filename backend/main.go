package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Límites de seguridad
	MaxPayloadSize = 52 * 1024 * 1024 // 52 MB para soportar imágenes de hasta 50MB + overhead
	MaxPINsPerConn = 5                // Máximo de PINs activos por conexión (evita saturación)
	MaxFailures    = 10               // Intentos máximos antes de desconectar
)

// ClipboardItem representa un dato guardado temporalmente (texto o binario).
type ClipboardItem struct {
	PayloadType string // "text" o "image"
	Text        string
	Binary      []byte
	Ext         string // Extensión para la imagen, ej: "png"
	Sender      *websocket.Conn
	ExpiresAt   time.Time
}

// Store maneja la memoria efímera de manera concurrente.
type Store struct {
	mu    sync.RWMutex
	items         map[string]ClipboardItem
	pinsCount     map[*websocket.Conn]int // Para limitar PINs por conexión
	failedAttempts map[*websocket.Conn]int // Control de fuerza bruta
}

func NewStore() *Store {
	s := &Store{
		items:          make(map[string]ClipboardItem),
		pinsCount:      make(map[*websocket.Conn]int),
		failedAttempts: make(map[*websocket.Conn]int),
	}
	go s.RunCleanup()
	return s
}

// GeneratePIN crea un código de 4 dígitos usando crypto/rand.
func (s *Store) GeneratePIN() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.generatePINUnsafe()
}

func (s *Store) generatePINUnsafe() string {
	for {
		// Generar número seguro entre 0 y 9999
		n, err := rand.Int(rand.Reader, big.NewInt(10000))
		if err != nil {
			log.Println("Error generando PIN seguro:", err)
			return "0000" // Fallback seguro
		}
		pin := fmt.Sprintf("%04d", n.Int64())
		if _, exists := s.items[pin]; !exists {
			return pin
		}
	}
}

// Save guarda el dato, controla el límite por conexión, y retorna el PIN.
func (s *Store) Save(item ClipboardItem) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.pinsCount[item.Sender] >= MaxPINsPerConn {
		return "", fmt.Errorf("límite de PINs alcanzado (max %d)", MaxPINsPerConn)
	}

	pin := s.generatePINUnsafe()
	item.ExpiresAt = time.Now().Add(5 * time.Minute)

	s.items[pin] = item
	s.pinsCount[item.Sender]++
	return pin, nil
}

// Retrieve obtiene el dato y lleva control de intentos fallidos.
func (s *Store) Retrieve(pin string, receiver *websocket.Conn) (ClipboardItem, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Control de fuerza bruta
	if s.failedAttempts[receiver] >= MaxFailures {
		return ClipboardItem{}, fmt.Errorf("demasiados intentos fallidos")
	}

	item, exists := s.items[pin]
	if !exists {
		s.failedAttempts[receiver]++
		return ClipboardItem{}, fmt.Errorf("PIN no encontrado o expirado")
	}
	
	// Éxito: borrar pin y resetear fallos (si aplicara)
	delete(s.items, pin)
	s.pinsCount[item.Sender]--
	return item, nil
}

// RemoveByConnection limpia datos cuando un cliente se desconecta.
func (s *Store) RemoveByConnection(conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for pin, item := range s.items {
		if item.Sender == conn {
			delete(s.items, pin)
		}
	}
	delete(s.pinsCount, conn)
	delete(s.failedAttempts, conn)
}

// RunCleanup es una goroutine que elimina pines expirados tras 5 minutos.
func (s *Store) RunCleanup() {
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		s.Cleanup()
	}
}

// Cleanup realiza una pasada de limpieza de pines expirados.
func (s *Store) Cleanup() {
	now := time.Now()
	s.mu.Lock()
	defer s.mu.Unlock()
	for pin, item := range s.items {
		if now.After(item.ExpiresAt) {
			delete(s.items, pin)
			if item.Sender != nil {
				if count, ok := s.pinsCount[item.Sender]; ok && count > 0 {
					s.pinsCount[item.Sender]--
				}
			}
		}
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Permitir cualquier origen temporalmente para que funcionen túneles como Ngrok
		return true
	},
}

type WSMessage struct {
	Action  string `json:"action"`
	Payload string `json:"payload,omitempty"`
	PIN     string `json:"pin,omitempty"`
	Ext     string `json:"ext,omitempty"`
}

func handleConnections(w http.ResponseWriter, r *http.Request, store *Store) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error al actualizar conexión a WebSocket:", err)
		return
	}
	defer conn.Close()
	defer store.RemoveByConnection(conn)

	// Límite de tamaño del payload para evitar DoS por memoria
	conn.SetReadLimit(MaxPayloadSize)

	var waitingForBinary bool
	var expectedExt string

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error de lectura WS: %v", err)
			}
			break
		}

		if messageType == websocket.BinaryMessage {
			if !waitingForBinary {
				continue // Ignorar binarios no esperados
			}
			pin, err := store.Save(ClipboardItem{
				PayloadType: "image",
				Binary:      p,
				Ext:         expectedExt,
				Sender:      conn,
			})
			waitingForBinary = false

			if err != nil {
				conn.WriteJSON(map[string]string{
					"action":  "error",
					"message": err.Error(),
				})
				continue
			}
			conn.WriteJSON(map[string]string{
				"action": "pin_generated",
				"pin":    pin,
			})
			continue
		}

		// Es un mensaje de texto (JSON)
		var msg WSMessage
		if err := json.Unmarshal(p, &msg); err != nil {
			continue // Omitir mensajes malformados
		}

		switch msg.Action {
		case "send_text":
			pin, err := store.Save(ClipboardItem{
				PayloadType: "text",
				Text:        msg.Payload,
				Sender:      conn,
			})
			if err != nil {
				conn.WriteJSON(map[string]string{
					"action":  "error",
					"message": err.Error(),
				})
				continue
			}
			conn.WriteJSON(map[string]string{
				"action": "pin_generated",
				"pin":    pin,
			})

		case "send_image":
			expectedExt = msg.Ext
			waitingForBinary = true
			conn.WriteJSON(map[string]string{
				"action": "ready_for_binary",
			})

		case "receive_text":
			item, err := store.Retrieve(msg.PIN, conn)
			if err != nil {
				conn.WriteJSON(map[string]string{
					"action":  "error",
					"message": err.Error(),
				})
				if err.Error() == "demasiados intentos fallidos" {
					return // Desconectar al cliente
				}
				continue
			}

			if item.PayloadType == "image" {
				// Enviar metadatos
				conn.WriteJSON(map[string]string{
					"action": "image_received",
					"ext":    item.Ext,
				})
				// Enviar binario
				conn.WriteMessage(websocket.BinaryMessage, item.Binary)
			} else {
				// Enviar texto
				conn.WriteJSON(map[string]string{
					"action":  "text_received",
					"payload": item.Text,
				})
			}

			// Notificar al emisor que la transferencia fue exitosa
			if item.Sender != nil {
				item.Sender.WriteJSON(map[string]string{
					"action": "transfer_complete",
				})
			}
		}
	}
}

func main() {
	store := NewStore()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleConnections(w, r, store)
	})

	log.Println("Servidor WebSocket backend corriendo en :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error en ListenAndServe:", err)
	}
}
