package src

import (
	"sync"
	"time"
)

// Entities
type Player struct {
	ID   string
	Name string
	Team *Team
}

type Team struct {
	ID      string
	Name    string
	Players []*Player
}

type Match struct {
	ID         string
	HomeTeam   *Team
	AwayTeam   *Team
	Date       time.Time
	Venue      string
	Status     MatchStatus
	Score      *Score
	Commentary []string
	mu         sync.RWMutex
}

type Score struct {
	HomeTeamRuns    int
	HomeTeamWickets int
	AwayTeamRuns    int
	AwayTeamWickets int
}

// Factories
func NewPlayer(name string, team *Team, id string) *Player {
	return &Player{
		ID:   id,
		Name: name,
		Team: team,
	}
}

func NewTeam(name string, id string) *Team {
	return &Team{
		ID:   id,
		Name: name,
	}
}

func NewMatch(homeTeam *Team, awayTeam *Team, date time.Time, venue string, id string) *Match {
	return &Match{
		ID:         id,
		HomeTeam:   homeTeam,
		AwayTeam:   awayTeam,
		Date:       date,
		Venue:      venue,
		Status:     Scheduled,
		Score:      &Score{},
		Commentary: []string{},
	}
}
