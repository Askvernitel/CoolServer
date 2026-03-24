package game

import (
	. "project_go/internal/game/interfaces"

	"github.com/gorilla/websocket"
)

type GameState struct {
	players []*Player
}

func (gs *GameState) Update(ops []OperationBundle) {
	gs.applyOperations(ops)
}

func (gs *GameState) applyOperations(ops []OperationBundle) {
	for _, op := range ops {
		gs.handleOperationBundle(op)

	}
}
func (gs *GameState) handleOperationBundle(opb OperationBundle) {
	for _, op := range opb.GetOperations() {
		gs.handleOperation(opb.GetCaller(), op)
	}
}
func (gs *GameState) handleOperation(caller OperationCaller, op Operation) {
	p := gs.getPlayerById(Id(caller))
	switch op.GetType() {
	case RUN_RIGHT_OPERATION:
		p.MoveRight()
	case RUN_LEFT_OPERATION:
		p.MoveLeft()
	case RUN_UP_OPERATION:
		p.MoveUp()
	case RUN_DOWN_OPERATION:
		p.MoveDown()
	}
}
func (gs *GameState) GetState() *GameStateData {
	return &GameStateData{
		Players: gs.players,
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

type GameStateData struct {
	Players []*Player `json:"player"`
}

func (gsd *GameStateData) GetData() *GameStateData {
	return &GameStateData{}
}

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
	Type   OperationType   `json:"type"`
	Caller OperationCaller `json:"caller"`
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

func (c *Conn) Write(s *GameStateData) {
	c.Conn.WriteJSON(s)
}

func (c *Conn) ReadOperations() []OperationBundle {
	return c.opBundles
}

func (c *Conn) Flush() {
	c.opBundles = nil
}

func (c *Conn) Read() {
	op := &GameOperationBundle{}
	c.Conn.ReadJSON(op)
	c.opBundles = append(c.opBundles, op)
}
