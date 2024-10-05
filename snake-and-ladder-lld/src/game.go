package src

import (
	"fmt"
	"sort"
)

type IGame interface {
	AddPlayer(player IGamePlayer) error
	NextMove() error
	GetGameStatus() GameStatus
	GetWinner() (int, error)
	StartGame() error
}

type Game struct {
	gameStatus       GameStatus
	currPlayer       int
	players          []IGamePlayer
	gameBoard        IGameBoard
	winner           IGamePlayer
	nextMoveStrategy INextMoveStrategy
}

func NewGame(gameBoard IGameBoard, nextMoveStrategy INextMoveStrategy) IGame {
	return &Game{
		gameStatus:       Created,
		currPlayer:       -1,
		players:          make([]IGamePlayer, 0),
		gameBoard:        gameBoard,
		winner:           nil,
		nextMoveStrategy: nextMoveStrategy,
	}
}
func (g *Game) StartGame() error {
	if g.gameStatus == Created {
		g.gameStatus = Started
		return nil
	}
	return ErrGameIsInInvalidState
}
func (g *Game) AddPlayer(player IGamePlayer) error {
	g.players = append(g.players, player)
	return nil
}

func (g *Game) NextMove() error {
	g.currPlayer = (g.currPlayer + 1) % len(g.players)
	nextMove := g.nextMoveStrategy.GetNextMove()
	player := g.players[g.currPlayer]
	userId, _ := player.GetUserId()
	curr, err := player.GetCurrentPosition()
	if err != nil {
		return err
	}
	fmt.Printf("######### player %d is moving, the current position is: %d \n", userId, curr)
	if g.gameBoard.IsPointOutside(curr + nextMove) {
		return nil
	}
	player.SetNewPosition(curr + nextMove)
	newCurr, err := player.GetCurrentPosition()
	if err != nil {
		return err
	}
	fmt.Printf("######### new position is: %d \n", newCurr)
	gamePieces, err := g.gameBoard.GetGamePieceAtStartPoint(newCurr)
	if err != nil {
		return err
	}
	if len(gamePieces) > 0 {
		err := g.handlePlayerMovementUsingGamePiece(gamePieces, player)
		if err != nil {
			return err
		}
	}
	endPoint, err := g.gameBoard.GetEndingPoint()
	if err != nil {
		return err
	}
	if newCurr == endPoint {
		g.gameStatus = Completed
		g.winner = player
	}
	return nil
}
func (g *Game) handlePlayerMovementUsingGamePiece(gamePieces []IGamePieces, player IGamePlayer) error {
	sort.Slice(gamePieces, func(i, j int) bool {
		firstPriority, err := gamePieces[i].Priority()
		if err != nil {
			return false
		}
		secondPriority, err := gamePieces[j].Priority()
		if err != nil {
			return false
		}
		return firstPriority > secondPriority
	})
	if len(gamePieces) > 0 {
		gamePieces[0].Move(player)
	}
	newCurr, err := player.GetCurrentPosition()
	if err != nil {
		return err
	}
	fmt.Printf("######### new position is: %d \n", newCurr)
	return nil
}
func (g *Game) GetGameStatus() GameStatus {
	return g.gameStatus
}
func (g *Game) GetWinner() (int, error) {
	if g.winner == nil {
		return 0, ErrGameIsInInvalidState
	}
	return g.winner.GetPlayerId()
}
