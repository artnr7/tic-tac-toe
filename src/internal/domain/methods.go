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
// упустим лучший вариант - ничью
//
// # Допустим проигрыш у нас по умолчанию передан
//
// Описание алгоритма:
// Вызываем
func minimax(b base, compSide, curSide uint8, w *uint8, v *vec) {
	var move vec

	for i := range b.field {
		for j := range b.field[i] {
			// ->
			if b.field[i][j] == e {
				move = vec{int8(i), int8(j)}
				curB := b
				curB.field[i][j] = curSide

				if win(&curB, curSide) && curSide == compSide && *w < vic {
					*w = vic
				}
				if *w == vic {
					return
				}
				// Что значит 3-curSide? curSide может быть 1 или 2
				minimax(curB, compSide, 3-curSide, w, v)
				if *w == vic {
					*v = move
					return
				}
			}
			// ->
		}
	}
	if *w < draw {
		*w = draw
		*v = move
	}
}

func drawornot(g GameSession, w *uint8, v *vec) bool {
	minimax(g.Base, g.compSide, g.compSide, w, v)
	if *w == draw {
		return true
	}
	return false
}

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

	// little optimization, always you should put your figure to the centre of
	// the field, 'cause this is the most powerful strategy
	if g.Base.field[1][1] == e && moveOrder == 0 {
		g.Base.field[1][1] = g.compSide
	}

	var v vec
	var w uint8

	minimax(g.Base, g.compSide, g.compSide, &w, &v)
}

func isFilledBlock(block uint8) bool {
	return block == x || block == o
}

func isItRightBlock(block uint8) bool {
	return block == e || isFilledBlock(block)
}

func basesBlocksEq(block1, block2 uint8) bool {
	return block1 == block2
}

// Return is taken move opposite
func isOpposideSideBlockMove(block, ownSide uint8) bool {
	return block != ownSide
}

func isFieldChanged(blocksCnt, oldBlocksCnt int8) bool {
	return blocksCnt == oldBlocksCnt+1
}

// GameChangeValidate Returns whether the game state has changed outside of
// acceptable game behavior
// true = all fine, acceptable behavior
// false = bad behavior, cheating
func (g *GameSession) GameChangeValidate() error {
	acceptMove := false
	g.Base.blocksCnt = 0

	for i := range g.Base.field {
		for j := range g.Base.field[i] {

			if !isItRightBlock(g.Base.field[i][j]) {
				return errors.New("wrong file format")
			}

			if isFilledBlock(g.Base.field[i][j]) {
				g.Base.blocksCnt++
			}

			// Если у нас один блок не совпадает в
			// поле не совпадает, то это нормально,
			// потому что возможно противник
			// сделал ход
			if !basesBlocksEq(g.Base.field[i][j], g.oldBase.field[i][j]) &&
				!isOpposideSideBlockMove(g.Base.field[i][j], g.compSide) {
				if acceptMove {
					return errors.New("wrong file format")
				}
				acceptMove = true
			}
		}
	}

	if !isFieldChanged(g.Base.blocksCnt, g.oldBase.blocksCnt) {
		return errors.New("field not changed")
	}

	return nil
}

func (g *GameSession) IsGameEnd() bool {
	var v vec
	var w uint8

	if win(&g.Base, g.compSide) || win(&g.Base, 3-g.compSide) ||
		drawornot(*g, &w, &v) {
		return true
	}
	return false
}
