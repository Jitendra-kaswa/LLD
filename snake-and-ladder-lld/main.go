package main

import (
	"fmt"
	"os"

	"snake-and-ladder.com/src"
)

func main() {
	nextMoveStrategy := src.NewDiceStrategy()

	gameBoard := src.NewGameBoard(1, 100)
	gameBoard.AddGamePiece(src.NewLadder(41, 79, 1))
	gameBoard.AddGamePiece(src.NewLadder(22, 58, 1))
	gameBoard.AddGamePiece(src.NewLadder(4, 56, 1))
	gameBoard.AddGamePiece(src.NewLadder(14, 55, 1))
	gameBoard.AddGamePiece(src.NewLadder(12, 50, 1))
	gameBoard.AddGamePiece(src.NewLadder(54, 88, 1))
	gameBoard.AddGamePiece(src.NewSnake(96, 42, 1))
	gameBoard.AddGamePiece(src.NewSnake(94, 71, 1))
	gameBoard.AddGamePiece(src.NewSnake(48, 16, 1))
	gameBoard.AddGamePiece(src.NewSnake(37, 3, 1))
	gameBoard.AddGamePiece(src.NewSnake(28, 10, 1))

	game := src.NewGame(gameBoard, nextMoveStrategy)
	game.AddPlayer(src.NewGamePlayer(1, 11))
	game.AddPlayer(src.NewGamePlayer(2, 22))
	game.AddPlayer(src.NewGamePlayer(3, 33))

	game.StartGame()

	for {
		err := game.NextMove()
		if err != nil {
			fmt.Printf("there is a error in the next move: %v \n", err)
			os.Exit(0)
		}
		gameStatus := game.GetGameStatus()
		if gameStatus == src.Completed {
			winner, err := game.GetWinner()
			if err != nil {
				fmt.Printf("error is getting winner : %v \n", err)
			}
			fmt.Printf("the game winner is : %d \n", winner)
			os.Exit(0)
		}
	}
}
