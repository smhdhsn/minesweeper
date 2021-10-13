package main

import (
	"fmt"
	"strconv"

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

	fmt.Println(play(board, bombs))
}

func play(board game.Board, bombs int) int {
	board.PlantBomb(bombs)

	for {
		board.Draw()

	getRow:
		row, _ := interaction.PrintInteractiveDialog("Enter row: ", interaction.GetInput)
		rowOff, err := strconv.ParseInt(row, 0, 10)
		if err != nil || int(rowOff) > len(board) {
			interaction.Println("Invalid row number!", content.Red)
			goto getRow
		}

	getColumn:
		col, _ := interaction.PrintInteractiveDialog("Enter column: ", interaction.GetInput)
		colOff, err := strconv.ParseInt(col, 0, 10)
		if err != nil || int(colOff) > len(board[0]) {
			interaction.Println("Invalid column number!", content.Red)
			goto getColumn
		}

		result := board[rowOff-1][colOff-1].Reveale()

		if result == -1 {
			board.Over()
			command.Clear()
			board.Draw()
			return -1
		} else if result == 999 {
			return 0
		}

		command.Clear()
	}
}
