package src

import (
	"errors"

	"github.com/google/uuid"
)

type IdGenerationStrategy interface {
	GenerateId() string
}

type ScoringStrategy interface {
	UpdateScore(match *Match, runs int, wickets int) error
}

type CommentaryStrategy interface {
	AddCommentary(match *Match, comment string) error
}

// Id generation strategy implementation
type IdGenerationUsingUUID struct{}

func NewIdGenerationUsingUUID() IdGenerationStrategy {
	return &IdGenerationUsingUUID{}
}

func (iguu IdGenerationUsingUUID) GenerateId() string {
	return uuid.New().String()
}

type StandardScoringStrategy struct{}

func NewStandardScoringStrategy() ScoringStrategy {
	return &StandardScoringStrategy{}
}

func (s *StandardScoringStrategy) UpdateScore(match *Match, runs int, wickets int) error {
	match.mu.Lock()
	defer match.mu.Unlock()

	if match.Status != Live {
		return errors.New("cannot update score for non-live match")
	}

	match.Score.HomeTeamRuns += runs
	match.Score.HomeTeamWickets += wickets
	return nil
}

type BasicCommentaryStrategy struct{}

func NewBasicCommentaryStrategy() CommentaryStrategy {
	return &BasicCommentaryStrategy{}
}

func (s *BasicCommentaryStrategy) AddCommentary(match *Match, comment string) error {
	match.mu.Lock()
	defer match.mu.Unlock()

	if match.Status != Live {
		return errors.New("cannot add commentary for non-live match")
	}

	match.Commentary = append(match.Commentary, comment)
	return nil
}
