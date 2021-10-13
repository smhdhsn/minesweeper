package game

import (
	"log"
	"math/rand"
	"time"

	"github.com/smhdhsn/minesweeper/command"
	"github.com/smhdhsn/minesweeper/content"
	"github.com/smhdhsn/minesweeper/interaction"
)

var difficulties map[string]map[string]int = map[string]map[string]int{
	"HARD":     {"width": 30, "height": 16, "bombs": 99},
	"MODERATE": {"width": 16, "height": 16, "bombs": 40},
	"EASY":     {"width": 9, "height": 9, "bombs": 10},
}

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

// generateRandom, generates a random number between given values.
func generateRandom(min, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min+1) + min
}
