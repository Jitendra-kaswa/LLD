package src

import (
	"time"

	"math/rand"
)

type INextMoveStrategy interface {
	GetNextMove() int
}

type DiceStrategy struct{}

func NewDiceStrategy() INextMoveStrategy {
	return &DiceStrategy{}
}

func (dc DiceStrategy) GetNextMove() int {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(6) + 1
	return randomInt
}
