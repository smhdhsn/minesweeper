package content

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

const (
	font   = "chunky"
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	White  = "\033[37m"
	Black  = "\033[30m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
)

// Banner creates a banner with given title and color.
func Banner(title string) string {
	return figure.NewFigure(title, font, true).String()
}

// DifficultyDialog creates the dialog related to difficulty with proper coloring and returns the created dialog as string.
func DifficultyDialog() string {
	return fmt.Sprintf("%vChoose a difficulty level:\n%v1) Hard     [16x30 - 99 bombs]\n%v2) Moderate [16x16 - 40 bombs]\n%v3) Easy     [ 9x9  - 10 bombs]%v\n", White, Red, Yellow, Green, Reset)
}
