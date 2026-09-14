package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestStore_GeneratePIN(t *testing.T) {
	s := NewStore()
	pin := s.GeneratePIN()
	if len(pin) != 4 {
		t.Errorf("Expected PIN length 4, got %d", len(pin))
	}
}

func TestStore_Save(t *testing.T) {
	s := NewStore()
	conn := &websocket.Conn{}
	item := ClipboardItem{
		PayloadType: "text",
		Text:        "hello",
		Sender:      conn,
	}

	pin, err := s.Save(item)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if pin == "" {
		t.Error("Expected non-empty PIN")
	}

	s.mu.Lock()
	count := s.pinsCount[conn]
	s.mu.Unlock()

	if count != 1 {
		t.Errorf("Expected pinsCount 1, got %d", count)
	}
}

func TestStore_SaveLimit(t *testing.T) {
	s := NewStore()
	conn := &websocket.Conn{}
	
	for i := 0; i < MaxPINsPerConn; i++ {
		_, err := s.Save(ClipboardItem{Sender: conn})
		if err != nil {
			t.Fatalf("Unexpected error at %d: %v", i, err)
		}
	}

	_, err := s.Save(ClipboardItem{Sender: conn})
	if err == nil {
		t.Error("Expected error due to MaxPINsPerConn, but got nil")
	}
}

func TestStore_Retrieve(t *testing.T) {
	s := NewStore()
	connSender := &websocket.Conn{}
	connReceiver := &websocket.Conn{}
	
	item := ClipboardItem{
		PayloadType: "text",
		Text:        "secret",
		Sender:      connSender,
	}

	pin, _ := s.Save(item)

	retrieved, err := s.Retrieve(pin, connReceiver)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if retrieved.Text != "secret" {
		t.Errorf("Expected 'secret', got '%s'", retrieved.Text)
	}

	// Should be deleted after retrieval
	_, err = s.Retrieve(pin, connReceiver)
	if err == nil {
		t.Error("Expected error retrieving same PIN twice")
	}
	
	s.mu.Lock()
	count := s.pinsCount[connSender]
	s.mu.Unlock()

	if count != 0 {
		t.Errorf("Expected pinsCount 0, got %d", count)
	}
}

func TestStore_RetrieveFailures(t *testing.T) {
	s := NewStore()
	conn := &websocket.Conn{}

	for i := 0; i < MaxFailures; i++ {
		_, err := s.Retrieve("wrong", conn)
		if err == nil {
			t.Fatalf("Expected error for wrong PIN at %d", i)
		}
	}

	_, err := s.Retrieve("any", conn)
	if err == nil || err.Error() != "demasiados intentos fallidos" {
		t.Errorf("Expected 'demasiados intentos fallidos' error, got %v", err)
	}
}

func TestStore_RemoveByConnection(t *testing.T) {
	s := NewStore()
	conn := &websocket.Conn{}
	
	s.Save(ClipboardItem{Sender: conn})
	s.Save(ClipboardItem{Sender: conn})
	
	s.mu.Lock()
	itemsLen := len(s.items)
	s.mu.Unlock()

	if itemsLen != 2 {
		t.Errorf("Expected 2 items, got %d", itemsLen)
	}

	s.RemoveByConnection(conn)
	
	s.mu.Lock()
	itemsLen = len(s.items)
	_, countExists := s.pinsCount[conn]
	s.mu.Unlock()

	if itemsLen != 0 {
		t.Errorf("Expected 0 items, got %d", itemsLen)
	}
	if countExists {
		t.Error("Expected pinsCount entry to be deleted")
	}
}

func TestStore_Expiration(t *testing.T) {
	s := NewStore()
	conn := &websocket.Conn{}
	
	// Manually inject an expired item
	s.mu.Lock()
	pin := "1234"
	s.items[pin] = ClipboardItem{
		Sender:    conn,
		ExpiresAt: time.Now().Add(-1 * time.Minute),
	}
	s.pinsCount[conn] = 1
	s.mu.Unlock()

	// Trigger manual cleanup
	s.Cleanup()

	s.mu.Lock()
	_, exists := s.items[pin]
	count := s.pinsCount[conn]
	s.mu.Unlock()

	if exists {
		t.Error("Expired item was not removed")
	}
	
	if count != 0 {
		t.Errorf("Expected pinsCount 0 after expiration, got %d", count)
	}
}

func TestStore_Retrieve_Expired(t *testing.T) {
	s := NewStore()
	conn := &websocket.Conn{}
	
	s.mu.Lock()
	pin := "9999"
	s.items[pin] = ClipboardItem{
		Sender:    conn,
		ExpiresAt: time.Now().Add(-1 * time.Second), // Expired 1s ago
	}
	s.pinsCount[conn] = 1
	s.mu.Unlock()

	_, err := s.Retrieve(pin, &websocket.Conn{})
	if err == nil {
		t.Error("Expected error when retrieving expired PIN, but got nil")
	}
}

func TestStore_ConcurrentRetrieve(t *testing.T) {
	s := NewStore()
	connSender := &websocket.Conn{}
	pin, _ := s.Save(ClipboardItem{Text: "shared", Sender: connSender})

	const numReceivers = 10
	results := make(chan error, numReceivers)
	var wg sync.WaitGroup

	for i := 0; i < numReceivers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := s.Retrieve(pin, &websocket.Conn{})
			results <- err
		}()
	}

	wg.Wait()
	close(results)

	successes := 0
	for err := range results {
		if err == nil {
			successes++
		}
	}

	if successes != 1 {
		t.Errorf("Expected exactly 1 success, got %d", successes)
	}
}

func TestWebSocketFlow(t *testing.T) {
	store := NewStore()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleConnections(w, r, store)
	}))
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

	// Client A (Sender)
	dialer := websocket.Dialer{}
	connA, _, err := dialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect A: %v", err)
	}
	defer connA.Close()

	// Client B (Receiver)
	connB, _, err := dialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect B: %v", err)
	}
	defer connB.Close()

	// A sends text
	err = connA.WriteJSON(WSMessage{
		Action:  "send_text",
		Payload: "integration test",
	})
	if err != nil {
		t.Fatalf("A failed to send text: %v", err)
	}

	// A receives PIN
	var msgA WSMessage
	err = connA.ReadJSON(&msgA)
	if err != nil {
		t.Fatalf("A failed to receive PIN: %v", err)
	}
	if msgA.Action != "pin_generated" || len(msgA.PIN) != 4 {
		t.Fatalf("Unexpected msg from A: %+v", msgA)
	}
	pin := msgA.PIN

	// B sends PIN
	err = connB.WriteJSON(WSMessage{
		Action: "receive_text",
		PIN:    pin,
	})
	if err != nil {
		t.Fatalf("B failed to send PIN: %v", err)
	}

	// B receives text
	var msgB WSMessage
	err = connB.ReadJSON(&msgB)
	if err != nil {
		t.Fatalf("B failed to receive text: %v", err)
	}
	if msgB.Action != "text_received" || msgB.Payload != "integration test" {
		t.Fatalf("Unexpected msg from B: %+v", msgB)
	}

	// A receives transfer_complete
	err = connA.SetReadDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		t.Fatalf("Failed to set deadline: %v", err)
	}
	err = connA.ReadJSON(&msgA)
	if err != nil {
		t.Fatalf("A failed to receive transfer_complete: %v", err)
	}
	if msgA.Action != "transfer_complete" {
		t.Fatalf("Expected transfer_complete, got: %+v", msgA)
	}
}

func TestWebSocketBinaryFlow(t *testing.T) {
	store := NewStore()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleConnections(w, r, store)
	}))
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

	// Client A (Sender)
	dialer := websocket.Dialer{}
	connA, _, err := dialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect A: %v", err)
	}
	defer connA.Close()

	// Client B (Receiver)
	connB, _, err := dialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect B: %v", err)
	}
	defer connB.Close()

	// A sends intent to send image
	err = connA.WriteJSON(WSMessage{
		Action:   "send_image",
		FileName: "test-image.png",
	})
	if err != nil {
		t.Fatalf("A failed to send intent: %v", err)
	}

	// A receives ready_for_binary
	var msgA WSMessage
	err = connA.ReadJSON(&msgA)
	if err != nil {
		t.Fatalf("A failed to receive ready_for_binary: %v", err)
	}
	if msgA.Action != "ready_for_binary" {
		t.Fatalf("Expected ready_for_binary, got: %+v", msgA)
	}

	// A sends binary data
	binaryData := []byte{0x89, 'P', 'N', 'G', 0x0D, 0x0A, 0x1A, 0x0A} // PNG header
	err = connA.WriteMessage(websocket.BinaryMessage, binaryData)
	if err != nil {
		t.Fatalf("A failed to send binary data: %v", err)
	}

	// A receives PIN
	err = connA.ReadJSON(&msgA)
	if err != nil {
		t.Fatalf("A failed to receive PIN: %v", err)
	}
	if msgA.Action != "pin_generated" {
		t.Fatalf("Expected pin_generated, got: %+v", msgA)
	}
	pin := msgA.PIN

	// B sends PIN
	err = connB.WriteJSON(WSMessage{
		Action: "receive_text",
		PIN:    pin,
	})
	if err != nil {
		t.Fatalf("B failed to send PIN: %v", err)
	}

	// B receives file_received
	var msgB WSMessage
	err = connB.ReadJSON(&msgB)
	if err != nil {
		t.Fatalf("B failed to receive file_received: %v", err)
	}
	if msgB.Action != "file_received" || msgB.FileName != "test-image.png" {
		t.Fatalf("Unexpected msg from B: %+v", msgB)
	}

	// B receives binary data
	msgType, p, err := connB.ReadMessage()
	if err != nil {
		t.Fatalf("B failed to receive binary data: %v", err)
	}
	if msgType != websocket.BinaryMessage {
		t.Fatalf("Expected BinaryMessage, got %d", msgType)
	}
	if string(p) != string(binaryData) {
		t.Errorf("Expected binary data %v, got %v", binaryData, p)
	}

	// A receives transfer_complete
	err = connA.SetReadDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		t.Fatalf("Failed to set deadline: %v", err)
	}
	err = connA.ReadJSON(&msgA)
	if err != nil {
		t.Fatalf("A failed to receive transfer_complete: %v", err)
	}
	if msgA.Action != "transfer_complete" {
		t.Fatalf("Expected transfer_complete, got: %+v", msgA)
	}
}

func TestMaxPayloadSize(t *testing.T) {
	store := NewStore()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleConnections(w, r, store)
	}))
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a payload larger than MaxPayloadSize
	largePayload := strings.Repeat("a", MaxPayloadSize+100)
	err = conn.WriteJSON(WSMessage{
		Action:  "send_text",
		Payload: largePayload,
	})
	
	// We might get an error on write or on subsequent read because the server closes the connection
	if err != nil {
		return // OK
	}

	// Wait for connection to close
	_, _, err = conn.ReadMessage()
	if err == nil {
		t.Error("Expected connection to close or error when sending payload exceeding MaxPayloadSize")
	}
}
