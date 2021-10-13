package game

type cell struct {
	areaBombs  int
	col, row   int
	isBomb     bool
	isDefused  bool
	isRevealed bool
	board      Board
}

// Reveal, reveals the cell's value.
// Returns (-1) if it's a bomb, and (areaBombs) if it's not a bomb.
func (c *cell) Reveale() int {
	if c.isDefused {
		return 0
	}

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

// Defuse, changes the 'isDefused' to opposite status of current status every time it's called.
func (c *cell) Defuse() {
	c.isDefused = !c.isDefused
}
