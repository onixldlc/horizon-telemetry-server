package handler

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// Manager maintains active connections
type Manager struct {
	Connections map[*websocket.Conn]bool
	lock        sync.Mutex
}

// NewManager creates a new WebSocket Manager
func NewManager() *Manager {
	return &Manager{
		Connections: make(map[*websocket.Conn]bool),
	}
}

// Register registers a new WebSocket connection
func (m *Manager) Register(conn *websocket.Conn) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.Connections[conn] = true
}

// Unregister removes a WebSocket connection
func (m *Manager) Unregister(conn *websocket.Conn) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.Connections, conn)
	conn.Close()
}

// SendMessage sends a message to all registered connections
func (m *Manager) SendMessage(msg []byte) {
	m.lock.Lock()
	defer m.lock.Unlock()
	for conn := range m.Connections {
		err := conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Printf("error: %v", err)
			m.Unregister(conn)
		}
	}
}
