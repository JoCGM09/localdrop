package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Límites de seguridad
	MaxPayloadSize = 10 * 1024 * 1024 // 10 MB para textos largos
	MaxPINsPerConn = 5                // Máximo de PINs activos por conexión (evita saturación)
	MaxFailures    = 10               // Intentos máximos antes de desconectar
)

// ClipboardItem representa un texto guardado temporalmente.
type ClipboardItem struct {
	Text      string
	Sender    *websocket.Conn
	ExpiresAt time.Time
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

// Save guarda el texto, controla el límite por conexión, y retorna el PIN.
func (s *Store) Save(text string, sender *websocket.Conn) (string, error) {
	s.mu.Lock()
	if s.pinsCount[sender] >= MaxPINsPerConn {
		s.mu.Unlock()
		return "", fmt.Errorf("límite de PINs alcanzado (max %d)", MaxPINsPerConn)
	}
	s.mu.Unlock()

	pin := s.GeneratePIN()
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[pin] = ClipboardItem{
		Text:      text,
		Sender:    sender,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
	s.pinsCount[sender]++
	return pin, nil
}

// Retrieve obtiene el texto y lleva control de intentos fallidos.
func (s *Store) Retrieve(pin string, receiver *websocket.Conn) (string, *websocket.Conn, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Control de fuerza bruta
	if s.failedAttempts[receiver] >= MaxFailures {
		return "", nil, fmt.Errorf("demasiados intentos fallidos")
	}

	item, exists := s.items[pin]
	if !exists {
		s.failedAttempts[receiver]++
		return "", nil, fmt.Errorf("PIN no encontrado o expirado")
	}
	
	// Éxito: borrar pin y resetear fallos (si aplicara)
	delete(s.items, pin)
	s.pinsCount[item.Sender]--
	return item.Text, item.Sender, nil
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
		now := time.Now()
		s.mu.Lock()
		for pin, item := range s.items {
			if now.After(item.ExpiresAt) {
				delete(s.items, pin)
			}
		}
		s.mu.Unlock()
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Seguridad CORS: Solo permitir orígenes de localhost o red local 192.168.x.x
		origin := r.Header.Get("Origin")
		if origin == "" {
			return true // Permitir clientes no-browser que no envían Origin (como curl o tools CLI)
		}
		// Validar (simplificado): localhost, 127.0.0.1, y red privada
		if strings.Contains(origin, "localhost") || strings.Contains(origin, "127.0.0.1") || strings.Contains(origin, "192.168.") {
			return true
		}
		log.Printf("Origen denegado: %s", origin)
		return false
	},
}

type WSMessage struct {
	Action  string `json:"action"`
	Payload string `json:"payload,omitempty"`
	PIN     string `json:"pin,omitempty"`
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

	for {
		var msg WSMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error de lectura WS: %v", err)
			}
			break
		}

		switch msg.Action {
		case "send_text":
			pin, err := store.Save(msg.Payload, conn)
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
		case "receive_text":
			text, senderConn, err := store.Retrieve(msg.PIN, conn)
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

			// Enviar el texto al receptor
			conn.WriteJSON(map[string]string{
				"action":  "text_received",
				"payload": text,
			})

			// Notificar al emisor que la transferencia fue exitosa
			if senderConn != nil {
				senderConn.WriteJSON(map[string]string{
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
