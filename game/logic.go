package game

import (
	"github.com/smhdhsn/minesweeper/command"
)

const (
	GameOver      = -1
	GameContinues = 0
	Win           = 1
)

var difficulties map[string]map[string]int = map[string]map[string]int{
	"HARD":     {"width": 30, "height": 16, "bombs": 99},
	"MODERATE": {"width": 16, "height": 16, "bombs": 40},
	"EASY":     {"width": 9, "height": 9, "bombs": 10},
}

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
