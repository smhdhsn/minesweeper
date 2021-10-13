package game

import (
	"strconv"

	"github.com/smhdhsn/minesweeper/command"
	"github.com/smhdhsn/minesweeper/content"
	"github.com/smhdhsn/minesweeper/interaction"
)

const (
	GameOver      = -1
	GameContinues = 0
	Win           = 1
)

// Execute, executes an action on board based on input.
// Returns (-1) if the game has been over with a loss, (0) if the game is still going, (1) if the game is over with a win.
func Execute(board Board, row, col, option int) int {
	var result int
	switch option {
	case 1:
		result = board[row-1][col-1].Reveale()
	case 2:
		board[row-1][col-1].Defuse()
	}

	if result == -1 {
		board.Over()
		command.Clear()
		board.Draw()
		return GameOver
	}

	if board.IsFullyDefused() {
		command.Clear()
		board.Draw()
		return Win
	}

	return GameContinues
}

// GetAction, gets row, column, option from player and returns the result.
func GetAction(board Board) (int, int, int) {
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

getOption:
	op, _ := interaction.PrintInteractiveDialog("Enter command[1:open, 2:defuse]: ", interaction.GetInput)
	option, err := strconv.ParseInt(op, 0, 10)
	if err != nil || option > 2 && option < 1 {
		interaction.Println("Invalid command!", content.Red)
		goto getOption
	}

	return int(rowOff), int(colOff), int(option)
}
