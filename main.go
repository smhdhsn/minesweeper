package main

import (
	"github.com/smhdhsn/minesweeper/command"
	"github.com/smhdhsn/minesweeper/content"
	"github.com/smhdhsn/minesweeper/game"
	"github.com/smhdhsn/minesweeper/interaction"
)

func main() {
	command.Clear()
	interaction.PrintDialog(content.Banner("MINE SWEEPER"), content.Green)
	interaction.PrintInteractiveDialog("Press enter to continue...", interaction.GetInput)

	settings := game.Initiate()

	board, bombs := game.Make(settings)
	interaction.PrintDialog("Starting the game...", content.Green)
	command.Clear()

	gameResult := play(board, bombs)

	if gameResult == game.Win {
		playerWon()
	} else {
		playerLost()
	}
}

// play, initiates the game process and returns the result.
func play(board game.Board, bombs int) int {
	board.PlantBomb(bombs)

	for {
		board.Draw()

		row, col, option := game.GetAction(board)

		result := game.Execute(board, row, col, option)

		if result == game.GameOver || result == game.Win {
			return result
		}

		command.Clear()
	}
}

// playerWon, will be executed if player won the game.
func playerWon() {
	interaction.PrintDialog(content.Banner("YOU WON"), content.Green)
}

// playerLost, will be executed if player lost the game.
func playerLost() {
	interaction.PrintDialog(content.Banner("YOU LOST"), content.Red)
}
