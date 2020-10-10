package game

import (
	"fmt"

	"github.com/crucialcarl/game_command_pattern/internal/command"
	"github.com/crucialcarl/game_command_pattern/internal/npc"
	"github.com/crucialcarl/game_command_pattern/internal/player"
)

type Game struct {
	player *player.Player
	npc    *npc.NPC
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Run() {
	g.player = player.NewPlayer()
	g.npc = npc.NewNPC()

	fmt.Println("player: ", g.player.GetRoom())
	fmt.Println("npc: ", g.npc.GetRoom())

	moveCommand := command.MoveCommand{g.player.Entity, 2}
	moveCommand2 := command.MoveCommand{g.npc.Entity, 1}

	tasks := []command.Command{
		moveCommand,
		moveCommand2,
	}

	executor := &command.Executor{
		Commands: tasks,
	}

	executor.ExecuteCommands()
	fmt.Println("player: ", g.player.GetRoom())
	fmt.Println("npc: ", g.npc.GetRoom())
}
