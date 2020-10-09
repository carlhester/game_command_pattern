package npc

import "github.com/crucialcarl/game_command_pattern/internal/entity"

type NPC struct {
	*entity.Entity
}

func NewNPC() *NPC {
	return &NPC{
		entity.NewEntity(2),
	}
}
