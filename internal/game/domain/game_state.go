package domain

import (
	. "project_go/internal/game/interfaces"
	"slices"
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

func (gs *GameState) AddPlayer(p *Player) {
	gs.players = append(gs.players, p)
}
func (gs *GameState) RemovePlayer(p *Player) {
	delIndex := slices.Index(gs.players, p)
	if delIndex == -1 {
		return
	}
	slices.Delete(gs.players, delIndex, delIndex+1)
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
	Id    Id      `json:"id"`
	X     float32 `json:"x"`
	Y     float32 `json:"y"`
	Speed float32 `json:"speed"`
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
	Players []*Player `json:"players"`
}

func (gsd *GameStateData) GetData() *GameStateData {
	return &GameStateData{}
}
