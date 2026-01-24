package domain

import (
	"errors"
)

func whoseMove(base *base) bool {
	var xes, oes int8
	for i := range base.field {
		for j := range i {
			if j == x {
				xes++
			} else {
				oes++
			}
		}
	}

	return xes > oes
}

func inc(winRow *uint8, el, side uint8) {
	if el == side {
		(*winRow)++
	}
}

func win(side uint8, base *base) bool {
	var hWinRow, vWinRow, d1WinRow, d2WinRow uint8

	for i := range base.field {

		hWinRow, vWinRow, d1WinRow, d2WinRow = 0, 0, 0, 0
		for j := range base.field[i] {
			inc(&hWinRow, base.field[i][j], side)
			inc(&vWinRow, base.field[j][i], side)
			if i == j {
				inc(&d1WinRow, base.field[i][j], side)
			}
			if i == 2-j {
				inc(&d1WinRow, base.field[i][j], side)
			}
		}

		// main statement
		if hWinRow == 3 || vWinRow == 3 || d1WinRow == 3 || d2WinRow == 3 {
			return true
		}

	}

	return false
}

func minimax(g *GameSession) {
	var me, enemy, side uint8
	var move vec

	for {
		move = vec{0, 0}
		if win(side, &g.Base) {
			return move
		}
	}
}

// TODO
// PutNextApologiseMove put computer prefer next move with more productivity
// with minimax strategy used
func (g *GameSession) PutNextApologiseMove() {
	// It is always more effective to place a figure in the center of a
	// field on the first move
	var moveOrder uint8

	for i := range g.Base.field {
		for j := range g.Base.field[i] {
			if g.Base.field[i][j] == e {
				continue
			}
			moveOrder++
		}
	}

	// little optimization, always you should put your figure to the center of
	// the field, 'cause this is the most powerful strategy
	if g.Base.field[1][1] == e && moveOrder == 0 {
		g.Base.field[1][1] = g.compSide
	}

	minimax(&g.Base)
}

func isFilledBlock(block int8) bool {
	return block == x || block == o
}

func isItRightBlock(block int8) bool {
	return block == e || isFilledBlock(block)
}

func basesBlocksEq(block1, block2 int8) bool {
	return block1 == block2
}

// Return is taken move opposite
func isOpposideSideBlockMove(block, ownSide int8) bool {
	return block != ownSide
}

func isFieldChanged(blocksCnt, oldBlocksCnt int8) bool {
	return blocksCnt == oldBlocksCnt+1
}

// TODO: doesn't work
// Returns whether the game state has changed outside of acceptable game
// behavior
// true = all fine, acceptable behavior
// false = bad behavior, cheating
func (g *GameSession) GameChangeValidate() error {
	// acceptMove := false

	for i := range g.Base.field {
		for j := range g.Base.field[i] {

			if !isItRightBlock(g.Base.field[i][j]) {
				return errors.New("wrong file format")
			}

			if isFilledBlock(g.Base.field[i][j]) {
				g.Base.blocksCnt++
			}

			if !basesBlocksEq(g.Base.field[i][j], g.oldBase.field[i][j]) &&
				!isOpposideSideBlockMove(g.Base.field[i][j], g.compSide) {
			}
		}
	}

	if !isFieldChanged(g.Base.blocksCnt, g.oldBase.blocksCnt) {
		return errors.New("field not changed")
	}

	return nil
}
