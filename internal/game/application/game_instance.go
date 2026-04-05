package application

import (
	. "project_go/internal/game/domain"
	. "project_go/internal/game/interfaces"
	"slices"
)

type GameInstance struct {
	Conns     []ReadWriteConn
	GameState *GameState
}

func (gi *GameInstance) AddConn(conn ReadWriteConn) {
	gi.Conns = append(gi.Conns, conn)

	//NOTE: Separate this
	gi.GameState.AddPlayer(&Player{
		Id:    Id(123),
		X:     0,
		Y:     0,
		Speed: 10,
	})
	gi.GameState.AddPlayer(&Player{
		Id:    Id(124),
		X:     0,
		Y:     0,
		Speed: 10,
	})
}

func (gi *GameInstance) RemoveConn(conn ReadWriteConn) {
	delIndex := slices.Index(gi.Conns, conn)
	if delIndex == -1 {
		return
	}
	gi.Conns = slices.Delete(gi.Conns, delIndex, delIndex+1)
}

func (gi *GameInstance) Update() {
	ops := gi.getOperations()

	gi.GameState.Update(ops)

	gi.UpdateClient()

}
func (gi *GameInstance) UpdateClient() {
	for _, conn := range gi.Conns {
		conn.Write(gi.GameState.GetState())
	}
}

func (gi *GameInstance) getOperations() []OperationBundle {
	ops := []OperationBundle{}

	for _, op := range gi.Conns {
		ops = append(ops, op.ReadOperations()...)
		op.Flush()
	}

	return ops
}
