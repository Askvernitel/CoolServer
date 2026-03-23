package game

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var globalGame = GameInstance{

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Adapter struct {
	t *Ticker
}

func NewAdapter() {

}

func (a *Adapter) Create(c *gin.Context) {
	c.Request.GetBody()
}

func (a *Adapter) Connect(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println("connection failed")
		return
	}

	NewConn(conn)

}

func (a *Adapter) Start() {
	a.t.start()
}

func (a *Adapter) Remove(c *gin.Context) {

}

func (a *Adapter) Update() {

}
