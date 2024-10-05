package src

import "errors"

type GameStatus int

const (
	Created GameStatus = iota
	Started
	Completed
)

var ErrGameIsInInvalidState = errors.New("game is in invalid state")
var ErrUserNotExists = errors.New("user doesn't exist")
