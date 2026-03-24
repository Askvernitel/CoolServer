package game

import (
	"log"
	. "project_go/internal/game/interfaces"
	"time"
)

type TPS int

type Ticker struct {
	games []*GameInstance
	tps   TPS
}

func NewTicker(tps TPS) *Ticker {
	return &Ticker{
		games: []*GameInstance{},
		tps:   tps,
	}
}

func (t *Ticker) start() {
	delta := time.Second / time.Duration(t.tps)
	ticker := time.NewTicker(delta)

	for range ticker.C {
		t.update()
	}
}

func (t *Ticker) update() {
	for _, game := range t.games {
		game.Update()
	}
}

func (t *Ticker) AddGameInstance(game *GameInstance) {
	t.games = append(t.games, game)
}

type GameInstance struct {
	conns []*Conn

	gameState *GameState
}

func (gi *GameInstance) AddConn(conn *Conn) {
	gi.conns = append(gi.conns, conn)

	//NOTE: Maybe separate this
	gi.gameState.AddPlayer(&Player{
		Id:    Id("Daniel"),
		X:     0,
		Y:     0,
		Speed: 200,
	})
}

func (gi *GameInstance) Update() {
	ops := gi.getOperations()
	log.Printf("Operations: %v\n", ops)
	gi.gameState.Update(ops)
	gi.WriteGameState()

}
func (gi *GameInstance) UpdateClient() {
	for _, conn := range gi.conns {
		conn.Write(gi.gameState.GetState())
	}
}
func (gi *GameInstance) getOperations() []OperationBundle {
	ops := []OperationBundle{}

	for _, op := range gi.conns {
		ops = append(ops, op.ReadOperations()...)
		op.Flush()
	}

	return ops
}

func (gi *GameInstance) WriteGameState() {
	for _, conn := range gi.conns {
		conn.Write(gi.gameState.GetState())
	}
}
