package main

import "github.com/crucialcarl/game_command_pattern/internal/game"

func main() {
	game := game.NewGame()
	game.Run()
}
