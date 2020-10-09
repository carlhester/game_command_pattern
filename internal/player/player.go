package player

import "github.com/crucialcarl/game_command_pattern/internal/entity"

type Player struct {
	*entity.Entity
}

func NewPlayer() *Player {
	return &Player{entity.NewEntity(1)}
}
