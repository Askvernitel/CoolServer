package game

import (
	. "project_go/internal/game/interfaces"

	"github.com/gorilla/websocket"
)

type GameState struct {
	players []*Player
}

func (gs *GameState) Update(ops []Operation) {
	gs.applyOperations(ops)
}

func (gs *GameState) applyOperations(ops []Operation) {
	for _, op := range ops {
		player := gs.getPlayerById(Id(op.GetCaller()))

		switch op.GetType() {
		case RUN_LEFT_OPERATION:
			player.MoveLeft()
		case RUN_RIGHT_OPERATION:
			player.MoveRight()
		case RUN_UP_OPERATION:
			player.MoveUp()
		case RUN_DOWN_OPERATION:
			player.MoveDown()
		}
	}
}

func (gs *GameState) getPlayerById(id Id) *Player {

	for _, p := range gs.players {
		if p.Id == id {
			return p
		}
	}
	return nil
}

type Player struct {
	Id    Id
	X     float32
	Y     float32
	Speed float32
}

func (p *Player) NewPlayer(x float32, y float32, speed float32) *Player {
	return &Player{
		X:     x,
		Y:     y,
		Speed: speed,
	}
}
func (p *Player) Write() {

}
func (p *Player) Read() {

}
func (p *Player) MoveRight() {
	p.X += p.Speed
}

func (p *Player) MoveLeft() {
	p.X -= p.Speed
}
func (p *Player) MoveUp() {
	p.Y += p.Speed
}
func (p *Player) MoveDown() {
	p.Y -= p.Speed
}

type GameOperation struct {
	t OperationType `json:"type"`
}

func (g *GameOperation) GetType() OperationType {
	return g.t
}
func (g *GameOperation) GetCaller() OperationCaller {

}

type Conn struct {
	Conn *websocket.Conn
	q    []Operation
}

func NewConn(conn *websocket.Conn) *Conn {
	return &Conn{Conn: conn}
}

func (c *Conn) WriteState() {
	//c.Conn.WriteJSON()
}

func (c *Conn) ReadOperations() []Operation {
	return c.q
}

func (c *Conn) Flush() {
	c.q = nil
}

func (c *Conn) Read() {
	op := &GameOperation{}
	c.Conn.ReadJSON(op)
	c.q = append(c.q, op)
}
