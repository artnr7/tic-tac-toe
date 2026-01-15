package domain

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

func minimax(base *base) {}

func (g *gameSession) GetNextApologiseMove() vec {
	// It is always more effective to place a figure in the center of a field on the first move
	if g.base.field[1][1] == 0 {
		return vec{1, 1}
	}

	v := vec{0, 0}
	return v
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

// Returns whether the game state has changed outside of acceptable game behavior
// true = all fine, acceptable behavior
// false = bad behavior, cheating
func (g *gameSession) GameChangeValidate() bool {
	// acceptMove := false

	for i := range g.base.field {
		for j := range g.base.field[i] {

			if !isItRightBlock(g.base.field[i][j]) {
				return false
			}

			if isFilledBlock(g.base.field[i][j]) {
				g.base.blocksCnt++
			}

			if !basesBlocksEq(g.base.field[i][j], g.oldBase.field[i][j]) && !isOpposideSideBlockMove(g.base.field[i][j], g.compSide) {
				// if acceptMove == true {
				// 	return false
				// }
				// acceptMove = true
			}
		}
	}

	if !isFieldChanged(g.base.blocksCnt, g.oldBase.blocksCnt) {
		return false
	}

	return true
}
