package game

import (
	"fmt"

	"github.com/crucialcarl/game_command_pattern/pkg/command"
	"github.com/crucialcarl/game_command_pattern/pkg/player"
)

type Game struct {
	player *player.Player
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Run() {
	g.player = &player.Player{}
	g.player.SetRoom(1)

	fmt.Println(g.player)

	moveCommand := command.MoveCommand{g.player, 2}

	tasks := []command.Command{
		moveCommand,
	}

	executor := &command.Executor{
		Commands: tasks,
	}

	executor.ExecuteCommands()
	fmt.Println(g.player)
}
