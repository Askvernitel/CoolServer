package domain

import (
	"errors"
	"io"
	"log"
	. "project_go/internal/game/interfaces"

	"github.com/gorilla/websocket"
)

type GameOperationBundle struct {
	Ops    []*GameOperation `json:"operations"`
	Caller OperationCaller  `json:"caller"`
}

func (gob *GameOperationBundle) GetOperations() []Operation {
	ops := []Operation{}
	for _, gop := range gob.Ops {
		ops = append(ops, gop)
	}

	return ops
}
func (gob *GameOperationBundle) GetCaller() OperationCaller {
	return gob.Caller
}

type GameOperation struct {
	Type OperationType `json:"type"`
}

func (g *GameOperation) GetType() OperationType {
	return g.Type
}

type Conn struct {
	Conn      *websocket.Conn
	opBundles []OperationBundle
}

func NewConn(conn *websocket.Conn) *Conn {
	return &Conn{Conn: conn}
}

func (c *Conn) Write(resp any) {
	c.Conn.WriteJSON(resp)
}

func (c *Conn) ReadOperations() []OperationBundle {

	return c.opBundles
}

func (c *Conn) Flush() {
	c.opBundles = []OperationBundle{}
}

func (c *Conn) Read() {
	for {
		op := &GameOperationBundle{}
		err := c.Conn.ReadJSON(op)
		if err != nil && !errors.Is(err, io.EOF) {
			log.Printf("Connection Read Error %v", err)
			break
		}

		c.opBundles = append(c.opBundles, op)
	}
}
