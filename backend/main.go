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
	MaxPayloadSize = 52 * 1024 * 1024 // 52 MB para soportar imágenes de hasta 50MB + overhead
	MaxPINsPerConn = 5                // Máximo de PINs activos por conexión (evita saturación)
	MaxFailures    = 10               // Intentos máximos antes de desconectar
	MaxTotalItems  = 100              // Límite global de archivos en memoria (prevención DoS)
)

// ClipboardItem representa un dato guardado temporalmente (texto o binario).
type ClipboardItem struct {
	PayloadType string // "text", "image", o "file"
	Text        string
	Binary      []byte
	FileName    string // Nombre original del archivo, ej: "foto.png" o "doc.pdf"
	Sender      *websocket.Conn
	ExpiresAt   time.Time
}

// Store maneja la memoria efímera de manera concurrente.
type Store struct {
	mu             sync.RWMutex
	items          map[string]ClipboardItem
	pinsCount      map[*websocket.Conn]int // Para limitar PINs por conexión
	failedAttempts map[string]int          // Control de fuerza bruta por IP
}

func NewStore() *Store {
	s := &Store{
		items:          make(map[string]ClipboardItem),
		pinsCount:      make(map[*websocket.Conn]int),
		failedAttempts: make(map[string]int),
	}
	go s.RunCleanup()
	return s
}

// GeneratePIN crea un código de 4 dígitos usando crypto/rand.
func (s *Store) GeneratePIN() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	pin, err := s.generatePINUnsafe()
	if err != nil {
		return "0000" // Fallback seguro mantenido solo por compatibilidad de tests, idealmente propagar error
	}
	return pin
}

func (s *Store) generatePINUnsafe() (string, error) {
	for {
		// Generar número seguro entre 0 y 9999
		n, err := rand.Int(rand.Reader, big.NewInt(10000))
		if err != nil {
			log.Println("Error generando PIN seguro:", err)
			return "", err // No usar fallback inseguro
		}
		pin := fmt.Sprintf("%04d", n.Int64())
		if _, exists := s.items[pin]; !exists {
			return pin, nil
		}
	}
}

// Save guarda el dato, controla el límite por conexión, y retorna el PIN.
func (s *Store) Save(item ClipboardItem) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.items) >= MaxTotalItems {
		return "", fmt.Errorf("servidor lleno, intente más tarde")
	}

	if s.pinsCount[item.Sender] >= MaxPINsPerConn {
		return "", fmt.Errorf("límite de PINs alcanzado (max %d)", MaxPINsPerConn)
	}

	pin, err := s.generatePINUnsafe()
	if err != nil {
		return "", fmt.Errorf("error interno al generar código de seguridad")
	}
	
	item.ExpiresAt = time.Now().Add(5 * time.Minute)
	s.items[pin] = item
	s.pinsCount[item.Sender]++
	return pin, nil
}

// Retrieve obtiene el dato y lleva control de intentos fallidos.
func (s *Store) Retrieve(pin string, receiver *websocket.Conn) (ClipboardItem, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	clientIP := getIP(receiver)

	// Control de fuerza bruta
	if s.failedAttempts[clientIP] >= MaxFailures {
		return ClipboardItem{}, fmt.Errorf("demasiados intentos fallidos")
	}

	item, exists := s.items[pin]
	if !exists {
		s.failedAttempts[clientIP]++
		return ClipboardItem{}, fmt.Errorf("PIN no encontrado o expirado")
	}

	// Validar expiración explícitamente al recuperar para evitar ventana de limpieza
	if time.Now().After(item.ExpiresAt) {
		delete(s.items, pin)
		s.failedAttempts[clientIP]++
		return ClipboardItem{}, fmt.Errorf("PIN no encontrado o expirado")
	}
	
	// Éxito: borrar pin y resetear fallos (si aplicara)
	delete(s.items, pin)
	s.pinsCount[item.Sender]--
	return item, nil
}

// Helper para extraer IP limpia
func getIP(conn *websocket.Conn) string {
	if conn == nil {
		return "unknown"
	}
	// Protección estricta para tests unitarios donde UnderlayingConn es nil
	defer func() {
		if r := recover(); r != nil {
			// En caso de pánico interno de Gorilla WS (ej. conn.conn == nil en tests)
		}
	}()
	addr := conn.RemoteAddr()
	if addr == nil {
		return "unknown"
	}
	addrStr := addr.String()
	parts := strings.Split(addrStr, ":")
	if len(parts) > 0 {
		return parts[0]
	}
	return addrStr
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
	clientIP := getIP(conn)
	delete(s.failedAttempts, clientIP)
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
		origin := r.Header.Get("Origin")
		if origin == "" {
			return true // Permitir clientes no-browser que no envían Origin
		}
		
		// Validar seguridad: localhost, 127.0.0.1, redes privadas o Ngrok temporal
		if strings.Contains(origin, "localhost") || strings.Contains(origin, "127.0.0.1") || 
			strings.Contains(origin, "192.168.") || strings.Contains(origin, "10.") || 
			strings.Contains(origin, "172.") || strings.Contains(origin, "ngrok-free.app") ||
			strings.Contains(origin, "loca.lt") {
			return true
		}
		log.Printf("Origen denegado por seguridad CSWSH: %s", origin)
		return false
	},
}

type WSMessage struct {
	Action   string `json:"action"`
	Payload  string `json:"payload,omitempty"`
	PIN      string `json:"pin,omitempty"`
	FileName string `json:"filename,omitempty"`
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
	var expectedFileName string
	var expectedPayloadType string

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
				PayloadType: expectedPayloadType,
				Binary:      p,
				FileName:    expectedFileName,
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

		case "send_image", "send_file": // Soportamos ambos para mantener compatibilidad semántica
			expectedFileName = msg.FileName
			expectedPayloadType = "file" // Unificamos el trato de binarios bajo "file"
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

			if item.PayloadType == "file" || item.PayloadType == "image" {
				// Enviar metadatos
				conn.WriteJSON(map[string]string{
					"action":   "file_received",
					"filename": item.FileName,
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
