package gamefunc

import (
	"fmt"
	"math/rand"
	"time"
)

// Структура для игры
type Game struct {
	board       [24]int // 24 позиции на доске
	PlayerHomes [2]int  // Количество выведенных шашек для каждого игрока
}

// Инициализация игры
func NewGame() *Game {
	game := &Game{
		PlayerHomes: [2]int{0, 0},
	}
	// Начальное размещение шашек
	game.board[0] = 15   // 15 шашек игрока-человека на позиции 1
	game.board[23] = -15 // 15 шашек компьютера на позиции 24
	return game
}

// Бросок кубиков
func rollDice() (int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1, rand.Intn(6) + 1
}

// Отрисовка доски
func (g *Game) PrintBoard() {
	fmt.Println("Текущая доска:")
	for i := 0; i < 12; i++ {
		fmt.Printf("%2d: %-2d ", i+1, g.board[i])
	}
	fmt.Printf("\n\n\n")
	for i := 12; i < 24; i++ {
		fmt.Printf("%2d: %-2d ", i+1, g.board[i])
	}
	fmt.Println("\n")
}

// Проверка возможности хода
func (g *Game) isValidMove(player int, from int, dice int) bool {
	if from < 0 || from >= 24 {
		return false
	}
	// Проверка на шашки игрока
	if player == 1 && g.board[from] <= 0 {
		return false
	}
	if player == 2 && g.board[from] >= 0 {
		return false
	}
	// Целевая позиция должна быть доступной
	target := from + dice
	if player == 2 {
		target = from - dice
	}
	if target < 0 || target >= 24 {
		return false
	}
	// Проверка, что позиция либо пуста, либо занята своими шашками
	if player == 1 && g.board[target] < 0 {
		return false
	}
	if player == 2 && g.board[target] > 0 {
		return false
	}
	return true
}

// Ввод хода от игрока
func (g *Game) getPlayerMove(dice int) int {
	var from int
	for {
		fmt.Printf("Введите номер позиции, с которой хотите сделать ход на %d шагов: ", dice)
		fmt.Scan(&from)
		from-- // Преобразование в индекс массива (от 0 до 23)

		if g.isValidMove(1, from, dice) {
			return from
		} else {
			fmt.Println("Неверный ход, попробуйте снова.")
		}
	}
}

// Ход компьютера
func (g *Game) computerMove(dice int) {
	for from := 23; from >= 0; from-- {
		if g.isValidMove(2, from, dice) {
			target := from - dice
			fmt.Printf("Компьютер делает ход с %d на %d\n", from+1, target+1)
			g.board[from]++
			if target < 0 {
				g.PlayerHomes[1]++
			} else {
				g.board[target]--
			}
			break
		}
	}
}

// Ход игрока
func (g *Game) PlayerMove() {
	dice1, dice2 := rollDice()
	fmt.Println("=========================================================================================")
	fmt.Printf("Вы бросили кубики: %d и %d\n", dice1, dice2)

	for _, dice := range []int{dice1, dice2} {
		g.PrintBoard()
		from := g.getPlayerMove(dice)
		target := from + dice

		// Вывод информации о ходе
		if target >= 24 {
			fmt.Println("=========================================================================================")
			fmt.Printf("Вы сделали ход с %d и вывели шашку.\n", from+1)
		} else {
			fmt.Println("=========================================================================================")
			fmt.Printf("Вы сделали ход с %d на %d.\n", from+1, target+1)
		}

		g.board[from]--
		if target >= 24 {
			g.PlayerHomes[0]++
		} else {
			g.board[target]++
		}
	}
}

// Ход компьютера
func (g *Game) ComputerTurn() {
	dice1, dice2 := rollDice()
	fmt.Println("=========================================================================================")
	fmt.Printf("Компьютер бросил кубики: %d и %d\n", dice1, dice2)

	for _, dice := range []int{dice1, dice2} {
		g.PrintBoard()
		for from := 23; from >= 0; from-- {
			if g.isValidMove(2, from, dice) {
				target := from - dice

				// Вывод информации о ходе
				if target < 0 {
					fmt.Println("=========================================================================================")

					fmt.Printf("Компьютер сделал ход с %d и вывел шашку.\n", from+1)
				} else {
					fmt.Println("=========================================================================================")
					fmt.Printf("Компьютер сделал ход с %d на %d.\n", from+1, target+1)
				}

				g.board[from]++
				if target < 0 {
					g.PlayerHomes[1]++
				} else {
					g.board[target]--
				}
				break
			}
		}
	}
}
