package service_impl

import (
	"domain"
	"errors"
	"fmt"
	"log"
	"math/rand"

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
			// fmt.Println(
			// 	"base = ",
			// 	base,
			// 	"\nhWinRow = ",
			// 	hWinRow,
			// 	"\nvWinRow = ",
			// 	vWinRow,
			// 	"\nd1WinRow = ",
			// 	d1WinRow,
			// 	"\nd2WinRow = ",
			// 	d2WinRow,
			// 	"\ni = ", i,
			// 	"\nj = ", j,
			// 	"\nbase.Field[i][j] = ", base.Field[i][j],
			// 	"\nbase.Field[j][i] = ", base.Field[j][i],
			// 	"\nside = ", side,
			// )
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

func fsd(gs *domain.GameSession) domain.Vec {
	var v domain.Vec
	var maxHeur int
	var fst bool = true

	for i := range gs.Base.Field {
		for j := range gs.Base.Field[i] {
			if isItEmptyBlock(gs.Base.Field[i][j]) {
				// ->
				curB := gs.Base
				curB.Field[i][j] = gs.CompSide
				heur := minimax(curB, gs.CompSide, gs.CompSide, -1)
				if fst {
					maxHeur = heur
					v = domain.Vec{int8(i), int8(j)}
					fst = false
				}

				if heur > maxHeur {
					maxHeur = heur
					v = domain.Vec{int8(i), int8(j)}
				}
				fmt.Println("***********")
				fmt.Println("HEUR = ", heur)
				fmt.Println("MAXHEUR = ", maxHeur)
				fmt.Println("v = ", v)
				fmt.Println("***********\n")

				// ->
			}
		}
	}
	return v
}

func minimax(b domain.Base, compSide, curSide uint8, value int) int {
	// fmt.Println("minimax")
	var heur, maxHeur, minHeur int
	var fst bool
	value++
	for i := range b.Field {
		for j := range b.Field[i] {
			if isItEmptyBlock(b.Field[i][j]) {
				// ->
				curB := b
				if value > 0 {
					curB.Field[i][j] = curSide
				}
				status := winOrDraw(&curB, curSide)
				switch status {
				case domain.Vic:
					if isItCompSide(compSide, curSide) {
						return 10 - value
					} else {
						return -10
					}
				case domain.Draw:
					return 0
				}

				heur = minimax(curB, compSide, 3-curSide, value)

				if fst {
					if isItCompSide(compSide, curSide) {
						maxHeur = heur
					} else {
						minHeur = heur
					}
					fst = false
				}

				if isItCompSide(compSide, curSide) {
					if heur > maxHeur {
						maxHeur = heur
					}
				} else {
					if heur < minHeur {
						minHeur = heur
					}
				}

				// ->
			}
		}
	}
	if isItCompSide(compSide, curSide) {
		return maxHeur
	} else {
		return minHeur
	}
}

// MakeNextMove put computer prefer next move with more productivity
// with minimax strategy used
func (g *ServiceImpl) MakeNextMove(gs *domain.GameSession) {
	log.Println(
		"********************************************************************make next move",
	)
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
		return
	}

	// minimax
	v := fsd(gs)

	gs.Base.Field[v.Y][v.X] = gs.CompSide
	gs.Base.BlocksCnt++
	log.Println("end make next move")
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
		fmt.Println(newGS.Base)
		fmt.Println(oldGS.Base)
		return errors.New("field not changed")
	}

	return nil
}

func (g *ServiceImpl) IsGameEnd(gs *domain.GameSession) error {
	fmt.Println("is game end")
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
	fmt.Println("end is game end")
	return nil
}
