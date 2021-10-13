package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Reader, reads input from command line.
var reader = bufio.NewReader(os.Stdin)

// GetInput reads a line of string from reader's input.
func GetInput() (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	input = strings.ReplaceAll(input, "\n", "")
	return input, nil
}
