package service_impl

import (
	"errors"
	"math/rand"

	"tictactoe/internal/domain"

	"github.com/google/uuid"
)

func inc(winRow *uint8, el, side uint8) {
	if el == side {
		(*winRow)++
	}
}

func winOrDraw(base *domain.Base, side uint8) int {
	var hWinRow, vWinRow, d1WinRow, d2WinRow, cnt uint8

	for i := range base.Field {

		hWinRow, vWinRow = 0, 0
		for j := range base.Field[i] {
			if !isItEmptyBlock(base.Field[i][j]) {
				cnt++
			}
			inc(&hWinRow, base.Field[i][j], side)
			inc(&vWinRow, base.Field[j][i], side)
			if i == j {
				inc(&d1WinRow, base.Field[i][j], side)
			}
			if i == 2-j {
				inc(&d2WinRow, base.Field[i][j], side)
			}
		}

		if hWinRow == 3 || vWinRow == 3 || d1WinRow == 3 || d2WinRow == 3 {
			return domain.Vic
		}

	}

	if cnt == 9 {
		return domain.Draw
	}

	return domain.Motive
}

func IsCompFirstMove(gs *domain.GameSession) bool {
	return rand.Int31n(2)+1 != int32(gs.CompSide)
}

func isItCompSide(compSide, curSide uint8) bool {
	return compSide == curSide
}

func minimax(
	b domain.Base,
	compSide, curSide uint8,
	value int,
) (int, domain.Vec) {
	var heur, maxHeur, minHeur int
	var fst bool = true
	var vec domain.Vec
	value++

	for i := range b.Field {
		for j := range b.Field[i] {
			if isItEmptyBlock(b.Field[i][j]) {
				// ->
				curB := b
				curB.Field[i][j] = curSide

				status := winOrDraw(&curB, curSide)
				switch status {
				case domain.Vic:
					if isItCompSide(compSide, curSide) {
						heur = 10 - value
					} else {
						heur = -10
					}
				case domain.Draw:
					heur = 0

				case domain.Motive:
					heur, _ = minimax(curB, compSide, 3-curSide, value)
				}

				if fst {
					if isItCompSide(compSide, curSide) {
						vec = domain.Vec{Y: int8(i), X: int8(j)}
						maxHeur = heur
					} else {
						vec = domain.Vec{Y: int8(i), X: int8(j)}
						minHeur = heur
					}
					fst = false
				}

				if isItCompSide(compSide, curSide) {
					if heur > maxHeur {
						vec = domain.Vec{Y: int8(i), X: int8(j)}
						maxHeur = heur
					}
				} else {
					if heur < minHeur {
						vec = domain.Vec{Y: int8(i), X: int8(j)}
						minHeur = heur
					}
				}
				// ->
			}
		}
	}
	if isItCompSide(compSide, curSide) {
		return maxHeur, vec
	} else {
		return minHeur, vec
	}
}

func makeFirstMoveInWholeGame(gs *domain.GameSession) bool {
	/* little optimization, always you should put your figure to the centre of the field, 'cause this is the most powerful strategy */
	var firstMoveInWholeGame bool = true
	for i := range gs.Base.Field {
		for j := range gs.Base.Field[i] {
			if gs.Base.Field[i][j] != domain.E {
				firstMoveInWholeGame = false
				break
			}
		}
	}
	if firstMoveInWholeGame {
		if IsCompFirstMove(gs) {
			gs.Base.Field[1][1] = gs.CompSide
			gs.Base.BlocksCnt = 1
		}
		return true
	}
	return false
}

// MakeNextMove put computer prefer next move with more productivity
// with minimax strategy used
func (g *ServiceImpl) MakeNextMove(gs *domain.GameSession) {
	g.IsGameEnd(gs)
	if gs.CompStatus != domain.Motive {
		return
	}

	if makeFirstMoveInWholeGame(gs) == true {
		return
	}

	// minimax
	_, bestMove := minimax(gs.Base, gs.CompSide, gs.CompSide, -1)

	gs.Base.Field[bestMove.Y][bestMove.X] = gs.CompSide
	gs.Base.BlocksCnt++
	g.IsGameEnd(gs)
}

func isFilledBlock(block uint8) bool {
	return block == domain.X || block == domain.O
}

func isItEmptyBlock(block uint8) bool {
	return block == domain.E
}

// isItRightBlock returns is block is empty or X or O
func isItRightBlock(block uint8) bool {
	return isItEmptyBlock(block) || isFilledBlock(block)
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

func (g *ServiceImpl) IsGameEnd(gs *domain.GameSession) error {
	oldGS, err := g.repo.GetModel(&(gs.UUID))
	if err != nil {
		return err
	}
	gs.CompSide = oldGS.CompSide
	if winOrDraw(&gs.Base, gs.CompSide) == domain.Vic {
		gs.CompStatus = domain.Vic
	} else if winOrDraw(&gs.Base, 3-gs.CompSide) == domain.Vic {
		gs.CompStatus = domain.Def
	} else if winOrDraw(&gs.Base, gs.CompSide) == domain.Draw {
		gs.CompStatus = domain.Draw
	} else {
		gs.CompStatus = domain.Motive
	}
	return nil
}
