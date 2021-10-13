package interaction

import (
	"fmt"

	"github.com/smhdhsn/minesweeper/content"
)

// NewLine, prints a new line.
func NewLine(str string) {
	fmt.Printf("%v\n%v", str, content.Reset)
}

// Print, prints a presented string.
func Print(str, option string) {
	fmt.Print(option+str, content.Reset)
}

// Println, prints a presented string with a new line.
func Println(str, option string) {
	fmt.Println(option+str, content.Reset)
}

// PrintDialog, prints a presented dialog.
func PrintDialog(dialog, option string) {
	fmt.Println(option+dialog, content.Reset)
}

// PrintInteractiveDialog, prints the dialog and asks user for an input related to that dialog.
func PrintInteractiveDialog(dialog string, getter func() (string, error)) (string, error) {
	fmt.Print(dialog)
	return getter()
}
