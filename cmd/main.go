package main

import (
	"backgammon/internal/gamefunc"
	"fmt"
)

func main() {
	fmt.Printf("Положительные числа это ваши фишки, отрицательные - ПК\n")
	game := gamefunc.NewGame()
	for game.PlayerHomes[0] < 15 && game.PlayerHomes[1] < 15 {
		game.PrintBoard()
		game.PlayerMove() // Ход игрока
		if game.PlayerHomes[0] >= 15 {
			fmt.Println("Поздравляем, вы победили!")
			break
		}
		game.ComputerTurn() // Ход компьютера
		if game.PlayerHomes[1] >= 15 {
			fmt.Println("Компьютер победил!")
			break
		}
	}
}
