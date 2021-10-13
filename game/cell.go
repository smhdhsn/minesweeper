package game

type cell struct {
	areaBombs  int
	col, row   int
	isBomb     bool
	isRevealed bool
	board      Board
}

// Reveal, reveals the cell's value.
// Returns (-1) if it's a bomb, and (areaBombs) if it's not a bomb.
func (c *cell) Reveale() int {
	if c.isBomb {
		return -1
	}

	if !c.isRevealed {
		c.isRevealed = true
		c.areaBombs = c.board.countAreaBombs(c)

		if c.areaBombs == 0 {
			c.board.revealeArea(c)
		}
	}

	return c.areaBombs
}

// IsRevealed, returns the value of 'isRevealed' property of 'cell' struct.
func (c *cell) IsRevealed() bool {
	return c.isRevealed
}
