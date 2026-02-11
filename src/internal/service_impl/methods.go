package service_impl

import (
	"domain"
	"errors"
	"math/rand"

	"github.com/google/uuid"
)

// func whoseMove(base *domain.Base) bool {
// 	var xes, oes int8
// 	for i := range base.Field {
// 		for j := range i {
// 			if j == domain.X {
// 				xes++
// 			} else {
// 				oes++
// 			}
// 		}
// 	}
//
// 	return xes > oes
// }

func inc(winRow *uint8, el, side uint8) {
	if el == side {
		(*winRow)++
	}
}

func win(base *domain.Base, side uint8) bool {
	var hWinRow, vWinRow, d1WinRow, d2WinRow uint8

	for i := range base.Field {

		hWinRow, vWinRow = 0, 0
		for j := range base.Field[i] {
			inc(&hWinRow, base.Field[i][j], side)
			inc(&vWinRow, base.Field[j][i], side)
			if i == j {
				inc(&d1WinRow, base.Field[i][j], side)
			}
			if i == 2-j {
				inc(&d2WinRow, base.Field[i][j], side)
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
func minimax(b domain.Base, compSide, curSide uint8, w *uint8, v *domain.Vec) {
	var move domain.Vec

	for i := range b.Field {
		for j := range b.Field[i] {
			// ->
			if b.Field[i][j] == domain.E {
				move = domain.Vec{int8(i), int8(j)}
				curB := b
				curB.Field[i][j] = curSide

				if win(&curB, curSide) && curSide == compSide &&
					*w < domain.Vic {
					*w = domain.Vic
				}
				if *w == domain.Vic {
					return
				}
				// Что значит 3-curSide? curSide может быть 1 или 2
				minimax(curB, compSide, 3-curSide, w, v)
				if *w == domain.Vic {
					*v = move
					return
				}
			}
			// ->
		}
	}
	if *w < domain.Draw {
		*w = domain.Draw
		*v = move
	}
}

func drawornot(g domain.GameSession) bool {
	var v *domain.Vec
	var w *uint8
	minimax(g.Base, g.CompSide, g.CompSide, w, v)

	if *w == domain.Draw {
		return true
	}

	return false
}

// MakeNextMove put computer prefer next move with more productivity
// with minimax strategy used
func (g *ServiceImpl) MakeNextMove(gs *domain.GameSession) {
	/* little optimization, always you should put your figure to the centre of the field, 'cause this is the most powerful strategy */
	var firstMove uint8
	for i := range gs.Base.Field {
		for j := range gs.Base.Field[i] {
			if gs.Base.Field[i][j] == domain.E {
				continue
			}
			firstMove++
		}
	}
	if firstMove == 0 {
		if rand.Int31n(2)+1 != int32(gs.CompSide) {
			return
		}
		gs.Base.Field[1][1] = gs.CompSide
	}

	// minimax
	var v domain.Vec
	var w uint8

	minimax(gs.Base, gs.CompSide, gs.CompSide, &w, &v)

	gs.Base.Field[v.Y][v.X] = gs.CompSide
	// fmt.Println(v)
}

func isFilledBlock(block uint8) bool {
	return block == domain.X || block == domain.O
}

// isItRightBlock returns is block is empty or X or O
func isItRightBlock(block uint8) bool {
	return block == domain.E || isFilledBlock(block)
}

// basesBlocksEq returns true if blocks is equal
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
func (g *ServiceImpl) GameChangeValidate(
	newGS *domain.GameSession,
	uuid *uuid.UUID,
) error {
	oldGS, _ := g.repo.GetModel(uuid)
	acceptMove := false

	for i := range newGS.Base.Field {
		for j := range newGS.Base.Field[i] {

			if !isItRightBlock(newGS.Base.Field[i][j]) {
				return errors.New("wrong file format")
			}

			if isFilledBlock(newGS.Base.Field[i][j]) {
				newGS.Base.BlocksCnt++
			}

			// Если у нас один блок не совпадает в
			// поле не совпадает, то это нормально,
			// потому что возможно противник
			// сделал ход
			if // если у нас блок нового поля не равно блоку старого
			!basesBlocksEq(newGS.Base.Field[i][j],
				oldGS.Base.Field[i][j],
			) &&
				// и эти блоки относятся к разным сторонам игры
				isOpposideSideBlockMove(
					newGS.Base.Field[i][j],
					oldGS.CompSide,
				) {
				// то это нормально, ставим себе заметку, да один блок изменился
				if // но если это не единичный случай
				acceptMove {
					// то возвращаем ошибку
					return errors.New("wrong file format")
				}
				acceptMove = true
			}
		}
	}

	if !isFieldChanged(newGS.Base.BlocksCnt, oldGS.Base.BlocksCnt) {
		return errors.New("field not changed")
	}

	return nil
}

func (g *ServiceImpl) IsGameEnd(gs *domain.GameSession) {
	if win(&gs.Base, gs.CompSide) {
		gs.Status = domain.Vic
	} else if win(&gs.Base, 3-gs.CompSide) {
		gs.Status = domain.Def
	} else if drawornot(*gs) {
		gs.Status = domain.Draw
	} else {
		gs.Status = domain.Motive
	}
}
