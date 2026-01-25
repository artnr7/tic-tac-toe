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

func win(base *base, side uint8) bool {
	var hWinRow, vWinRow, d1WinRow, d2WinRow uint8

	for i := range base.field {

		hWinRow, vWinRow = 0, 0
		for j := range base.field[i] {
			inc(&hWinRow, base.field[i][j], side)
			inc(&vWinRow, base.field[j][i], side)
			if i == j {
				inc(&d1WinRow, base.field[i][j], side)
			}
			if i == 2-j {
				inc(&d2WinRow, base.field[i][j], side)
			}
		}

		// main statement
		if hWinRow == 3 || vWinRow == 3 || d1WinRow == 3 || d2WinRow == 3 {
			return true
		}

	}

	return false
}

// Главный принцип Минимакса – это
// нахождение варинта событий, в котором
// возможно получить максимальную выгоду.
// Градация успеха -> выигрыш -> ничья -> проигрыш
// Т.к. значения дискретны и ограничены, то
// алгоритм будет подразумевать только поиск
// выигрыша, начиная с самого худшего
// варианта – проигрыша, в частности порядок
// поиска будет следующим
// проигрыш -> ничья -> выигрыш
// Если выигрыш найден, то дальнейший поиск не имеет смысла.
// Если все варианты рассмотрены возвращается самый лучший вариант исхода
//
// Показатель ничьи -> заняты все клетки, и не
// выполнено ни одно из условий выигрыша,
// независимо от стороны
//
// Почему нет возможности возвращать только выигрыш?
// Потому что если в ходе поиска была ничья и
// мы её отбросили, а дальше будут
// встречаться только поражения, то мы
// упустим лучший вариант
//
// Допустим проигрыш у нас по умолчанию передан
func minimax(b base, compSide, curSide uint8, w *uint8) {
	var move vec

	for i := range b.field {
		for j := range b.field[i] {
			// ->
			if b.field[i][j] == e {
				move = vec{int8(i), int8(j)}
				curB := b
				curB.field[i][j] = curSide

				if win(&curB, curSide) {
					if curSide == compSide {
						*w = vic
				}

				minimax(curB, compSide, 3-curSide, w)
				if *w == vic {
					return
				}
			}
			// ->
		}
	}
	*w = draw
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

	minimax(g, g.compSide)
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

func (g *GameSession) IsGameEnd() bool {
	return true
}
