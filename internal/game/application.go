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
	err := append(t.games, game)

	if err != nil {
		log.Println("Error in Ticker")
		return
	}
}

type GameInstance struct {
	conns []*Conn

	s State
}

func (gi *GameInstance) Update() {
	ops := gi.getOperations()
	gi.s.Update(ops)
}

func (gi *GameInstance) getOperations() []Operation {
	ops := []Operation{}

	for _, op := range gi.conns {
		ops = append(ops, op.ReadOperations()...)
		op.Flush()
	}

	return ops
}
