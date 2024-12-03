package connectionmanager

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

type ConnectionManager struct {
	connections map[string]*websocket.Conn // userId -> connection
	mu          sync.RWMutex               // for safe channels
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[string]*websocket.Conn),
	}
}

func (cm *ConnectionManager) AddConnection(userId string, conn *websocket.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.connections[userId] = conn
}

func (cm *ConnectionManager) RemoveConnection(userId string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.connections, userId)
}

func (cm *ConnectionManager) GetConnection(userId string) (*websocket.Conn, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	conn, exists := cm.connections[userId]
	return conn, exists
}

func (cm *ConnectionManager) AllConnections() {
	for name := range cm.connections {
		fmt.Println(name, " ", "true")
	}
}
