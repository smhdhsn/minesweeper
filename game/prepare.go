package game

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/smhdhsn/minesweeper/command"
	"github.com/smhdhsn/minesweeper/content"
	"github.com/smhdhsn/minesweeper/interaction"
)

// Initiate, initiates the process of making the game by asking user for difficulty level.
func Initiate() map[string]int {
	command.Clear()
	interaction.PrintDialog(content.DifficultyDialog(), content.Reset)

askForInput:
	difficulty, err := interaction.PrintInteractiveDialog("Your choice: ", interaction.GetInput)
	if err != nil {
		log.Fatalf("can't continue: %s", err)
	}

	switch difficulty {
	case "1":
		return difficulties["HARD"]
	case "2":
		return difficulties["MODERATE"]
	case "3":
		return difficulties["EASY"]
	default:
		interaction.PrintDialog("Invalid input!", content.Red)
		goto askForInput
	}
}

// Make, makes the game with Initiate() function's provided settings.
func Make(settings map[string]int) (Board, int) {
	board := make(Board, settings["height"])
	for row := 0; row < len(board); row++ {
		board[row] = make([]cell, settings["width"])
		for col := 0; col < len(board[row]); col++ {
			board[row][col].row = row
			board[row][col].col = col
			board[row][col].board = board
		}
	}

	return board, settings["bombs"]
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
	if err != nil || option > 2 || option < 1 {
		interaction.Println("Invalid command!", content.Red)
		goto getOption
	}

	return int(rowOff), int(colOff), int(option)
}

// generateRandom, generates a random number between given values.
func generateRandom(min, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min+1) + min
}
