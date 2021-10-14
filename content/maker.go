package content

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

const (
	font = "chunky"

	Reset  = "\033[0m"
	Italic = "\033[3m"

	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
	Black  = "\033[30m"

	BackGroundBlack  = "\033[40m"
	BackGroundRed    = "\033[41m"
	BackGroundGreen  = "\033[42m"
	BackGroundYellow = "\033[43m"
	BackGroundBlue   = "\033[44m"
	BackGroundPurple = "\033[45m"
	BackGroundCyan   = "\033[46m"
	BackGroundWhite  = "\033[47m"
)

// Banner creates a banner with given title and color.
func Banner(title string) string {
	return figure.NewFigure(title, font, true).String()
}

// DifficultyDialog creates the dialog related to difficulty with proper coloring and returns the created dialog as string.
func DifficultyDialog() string {
	return fmt.Sprintf("%vChoose a difficulty level:\n%v1) Hard     [16x30 - 99 bombs]\n%v2) Moderate [16x16 - 40 bombs]\n%v3) Easy     [ 9x9  - 10 bombs]%v\n", White, Red, Yellow, Green, Reset)
}

// Line, draws a line with given string.
func Line(str string, count, index int) string {
	color := Blue
	spacing := " "
	var endLine string
	var startLine string

	if str == "█" {
		endLine += Blue + "▎" + Reset
	}

	if index != 0 {
		color = White
		spacing = ""
		startLine = "▕"
	}

	return fmt.Sprintf("%v%v%v%v%v%v%v%v\n", spacing, Blue, startLine, color, strings.Repeat(str, count), str, Reset, endLine)
}

func RevealedCell(color, areaBombs string, col int) string {
	str := " "

	if col > 0 {
		str = "█ "
	}

	// interaction.Print("█ ", content.White)
	// interaction.Print(areaBombs, color)

	return fmt.Sprintf("%v%v%v%v%v", White, str, areaBombs, color, Reset)
}

// UnrevealedCell, draws a cell that has not been revealed.
func UnrevealedCell(color string, col int) string {
	var str string

	if col > 0 {
		str = "█"
	}

	return fmt.Sprintf("%v%v%v▓▓%v", White, str, color, Reset)
}

// NumberLine, draws a line with numbers inside.
func NumberLine(count int, background string) string {
	var str string

	for i := 1; i <= count; i++ {
		str += " " + strconv.FormatInt(int64(i), 10)
		if i <= 9 {
			str += " "
		}
	}

	return fmt.Sprintf("▕%v %v%v  %v%v▎\n", BackGroundWhite, Italic, str, Reset, Blue)
}
