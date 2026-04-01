package application

import (
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

func (t *Ticker) Start() {
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
