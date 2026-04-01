package adapter

import (
	"log"
	"net/http"
	. "project_go/internal/game/application"
	. "project_go/internal/game/domain"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var globalGame = &GameInstance{
	Conns:     []*Conn{},
	GameState: &GameState{},
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type HttpAdapter struct {
	ticker *Ticker
}

func NewAdapter(ticker *Ticker) *HttpAdapter {
	return &HttpAdapter{ticker: ticker}
}

func (a *HttpAdapter) Create(c *gin.Context) {
	c.Request.GetBody()

}

func (a *HttpAdapter) Connect(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println("connection failed", err)
		return
	}

	upConn := NewConn(conn)
	globalGame.AddConn(upConn)
	go upConn.Read()
}

func (a *HttpAdapter) Start() {
	a.ticker.AddGameInstance(globalGame)
	a.ticker.Start()
}

func (a *HttpAdapter) Remove(c *gin.Context) {

}

func (a *HttpAdapter) Update() {

}
