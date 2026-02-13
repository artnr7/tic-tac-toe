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

func fsd(gs *domain.GameSession) domain.Vec {
	fmt.Println("---------------------- FSD")
	var (
		cnt      int8
		maxScore int16
		fst      bool = true
		move     domain.Vec
	)
	scores := []int16{}

	for i := range gs.Base.Field {
		for j := range gs.Base.Field[i] {
			if gs.Base.Field[i][j] == domain.E {
				curGS := *gs
				curGS.Base.Field[i][j] = curGS.CompSide
				// ->
				scores = append(scores, 0)
				minimax(curGS.Base, gs.CompSide, gs.CompSide, &scores[cnt], -1)
				if fst {
					maxScore = scores[cnt]
					move = domain.Vec{int8(i), int8(j)}
					fst = false
				}
				if scores[cnt] > maxScore {
					maxScore = scores[cnt]
					move = domain.Vec{int8(i), int8(j)}
				}
				fmt.Println("************")
				fmt.Println("i = ", i)
				fmt.Println("j = ", j)
				fmt.Println("SCORE = ", scores[cnt])
				fmt.Println("MAXSCORE = ", maxScore)
				fmt.Println("MOVE = ", move)
				fmt.Println("************")
				cnt++
				// ->
			}
		}
	}
	return move
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
func minimax(
	b domain.Base,
	compSide, curSide uint8,
	score *int16,
	value int16,
) {
	value++
	// fmt.Println("========================================== minimax")
	for i := range b.Field {
		for j := range b.Field[i] {
			if b.Field[i][j] == domain.E {
				// ->
				curB := b
				curB.Field[i][j] = curSide

				// fmt.Println(
				// 	"************",
				// 	"\ni = ", i, "\nj = ", j,
				// 	"\ncurb = \n",
				// 	curB.Field[0],
				// 	"\n",
				// 	curB.Field[1],
				// 	"\n",
				// 	curB.Field[2],
				// 	"\ncurside = ",
				// 	curSide,
				// 	"\ncompside = ",
				// 	compSide,
				// )

				status := winOrDraw(&curB, curSide)
				// fmt.Println("status is = ", status)
				if status == domain.Vic {
					if curSide == compSide {
						*score += 10 + value
					} else {
						*score -= 10 - value
					}
				}
				if status == domain.Vic || status == domain.Draw {
					// fmt.Println("--- score is ", *score)
					return
				}
				// fmt.Println("**********************")

				// Что значит 3-curSide? curSide может быть 1 или 2
				minimax(curB, compSide, 3-curSide, score, value)
				// ->
			}
		}
	}
	// fmt.Println(
	// 	"************",
	// 	"\ncurb = \n",
	// 	curB.Field[0],
	// 	"\n",
	// 	curB.Field[1],
	// 	"\n",
	// 	curB.Field[2],
	// 	"\ncurside = ",
	// 	curSide,
	// 	"\ncompside = ",
	// 	compSide,
	// 	"\nv = ",
	// 	*v,
	// 	"\nw = ",
	// 	*w,
	// 	"**************",
	// )
}

func IsCompFirstMove(gs *domain.GameSession) bool {
	return rand.Int31n(2)+1 != int32(gs.CompSide)
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
