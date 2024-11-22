package client

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID     string
	Conn   *websocket.Conn
	ChatID uuid.UUID
}
