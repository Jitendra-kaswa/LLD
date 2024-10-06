package src

import (
	"errors"

	"github.com/google/uuid"
)

// all strategies
type SplitStrategy interface {
	Split(expense *Expense) ([]*Split, error)
}

type IdGenerationStrategy interface {
	GenerateId() string
}

type BalanceStrategy interface {
	CalculateBalances(expenses []*Expense) (map[string]float64, error)
}

// Id generation strategy implementation
type IdGenerationUsingUUID struct{}

func NewIdGenerationUsingUUID() IdGenerationStrategy {
	return &IdGenerationUsingUUID{}
}

func (iguu IdGenerationUsingUUID) GenerateId() string {
	return uuid.New().String()
}

// Expense Split strategies implementation
type EqualSplitStrategy struct{}

func NewEqualSplitStrategy() SplitStrategy {
	return &EqualSplitStrategy{}
}

func (s *EqualSplitStrategy) Split(expense *Expense) ([]*Split, error) {
	if len(expense.Participants) == 0 {
		return nil, errors.New("no participants for split")
	}

	amountPerPerson := expense.Amount / float64(len(expense.Participants))
	splits := make([]*Split, len(expense.Participants))

	for i, participant := range expense.Participants {
		splits[i] = NewSplit(participant, amountPerPerson)
	}

	return splits, nil
}

type SimpleBalanceCalculationStrategy struct{}

func NewSimpleBalanceCalculationStrategy() BalanceStrategy {
	return &SimpleBalanceCalculationStrategy{}
}

func (s *SimpleBalanceCalculationStrategy) CalculateBalances(expenses []*Expense) (map[string]float64, error) {
	balances := make(map[string]float64)
	group := expenses[0].Group
	for _, member := range group.Members {
		balances[member.ID] = 0
	}

	for _, expense := range expenses {
		if expense.Group != nil && expense.Group.ID == group.ID {
			balances[expense.Payer.ID] += expense.Amount
			for _, split := range expense.Splits {
				balances[split.User.ID] -= split.Amount
			}
		}
	}

	return balances, nil
}
