package command

import "github.com/crucialcarl/game_command_pattern/pkg/player"

type MoveCommand struct {
	Player          *player.Player
	DestinationRoom int
}

func (m MoveCommand) Move(p *player.Player, d int) {
	m.Player = p
	m.DestinationRoom = d
}

func (m MoveCommand) Execute() {
	m.Player.SetRoom(m.DestinationRoom)
}
