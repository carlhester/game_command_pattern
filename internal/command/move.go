package command

import "github.com/crucialcarl/game_command_pattern/internal/entity"

type MoveCommand struct {
	Entity          *entity.Entity
	DestinationRoom int
}

func (m MoveCommand) Move(e *entity.Entity, d int) {
	m.Entity = e
	m.DestinationRoom = d
}

func (m MoveCommand) Execute() {
	m.Entity.SetRoom(m.DestinationRoom)
}
