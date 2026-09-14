package main

import (
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
