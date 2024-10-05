package src

import (
	"errors"
	"fmt"
)

// Entities

// Snake, Ladder, Player, Game, Gameboard
type IGamePieces interface {
	Move(user IGamePlayer) error
	GetStartPosition() int
	Priority() (int, error)
}

type Snake struct {
	startPosition int
	endPosition   int
	priority      int
}

func NewSnake(startPos, endPos, priority int) IGamePieces {
	return &Snake{
		startPosition: startPos,
		endPosition:   endPos,
		priority:      priority,
	}
}

func (s Snake) Move(user IGamePlayer) error {
	fmt.Printf("this is a snake from position %d to %d \n", s.startPosition, s.endPosition)
	position, err := user.GetCurrentPosition()
	if err != nil {
		return err
	}
	if position != s.startPosition {
		return errors.New("call to invalid snake")
	}
	user.SetNewPosition(s.endPosition)
	return nil
}
func (s Snake) Priority() (int, error) {
	return s.priority, nil
}

func (s Snake) GetStartPosition() int {
	return s.startPosition
}

type Ladder struct {
	startPosition int
	endPosition   int
	priority      int
}

func NewLadder(startPos, endPos, priority int) IGamePieces {
	return &Ladder{
		startPosition: startPos,
		endPosition:   endPos,
		priority:      priority,
	}
}

func (s Ladder) Move(user IGamePlayer) error {
	fmt.Printf("this is a ladder from position %d to %d \n", s.startPosition, s.endPosition)
	position, err := user.GetCurrentPosition()
	if err != nil {
		return err
	}
	if position != s.startPosition {
		return errors.New("call to invalid snake")
	}
	user.SetNewPosition(s.endPosition)
	return nil
}
func (s Ladder) Priority() (int, error) {
	return s.priority, nil
}

func (s Ladder) GetStartPosition() int {
	return s.startPosition
}
