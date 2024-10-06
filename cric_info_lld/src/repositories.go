package src

import (
	"fmt"
	"sync"
)

type MatchRepository interface {
	Save(match *Match) error
	FindByID(id string) (*Match, error)
	FindAll() ([]*Match, error)
	Update(match *Match) error
	Delete(id string) error
}

type TeamRepository interface {
	Save(team *Team) error
	FindByID(id string) (*Team, error)
	FindAll() ([]*Team, error)
	Update(team *Team) error
	Delete(id string) error
}

type PlayerRepository interface {
	Save(player *Player) error
	FindByID(id string) (*Player, error)
	FindAll() ([]*Player, error)
	Update(player *Player) error
	Delete(id string) error
}

type InMemoryMatchRepository struct {
	matches map[string]*Match
	mu      sync.RWMutex
}

func NewInMemoryMatchRepository() MatchRepository {
	return &InMemoryMatchRepository{
		matches: make(map[string]*Match),
	}
}

func (r *InMemoryMatchRepository) Save(match *Match) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.matches[match.ID] = match
	return nil
}

func (r *InMemoryMatchRepository) FindByID(id string) (*Match, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	match, ok := r.matches[id]
	if !ok {
		return nil, fmt.Errorf("match with ID %s not found", id)
	}
	return match, nil
}

func (r *InMemoryMatchRepository) FindAll() ([]*Match, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	matches := make([]*Match, 0, len(r.matches))
	for _, match := range r.matches {
		matches = append(matches, match)
	}
	return matches, nil
}

func (r *InMemoryMatchRepository) Update(match *Match) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.matches[match.ID] = match
	return nil
}

func (r *InMemoryMatchRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.matches, id)
	return nil
}

type InMemoryTeamRepository struct {
	teams map[string]*Team
	mu    sync.RWMutex
}

func NewInMemoryTeamRepository() TeamRepository {
	return &InMemoryTeamRepository{
		teams: make(map[string]*Team),
	}
}

func (r *InMemoryTeamRepository) Save(team *Team) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.teams[team.ID] = team
	return nil
}

func (r *InMemoryTeamRepository) FindByID(id string) (*Team, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	team, ok := r.teams[id]
	if !ok {
		return nil, fmt.Errorf("team with ID %s not found", id)
	}
	return team, nil
}

func (r *InMemoryTeamRepository) FindAll() ([]*Team, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	teams := make([]*Team, 0, len(r.teams))
	for _, team := range r.teams {
		teams = append(teams, team)
	}
	return teams, nil
}

func (r *InMemoryTeamRepository) Update(team *Team) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.teams[team.ID] = team
	return nil
}

func (r *InMemoryTeamRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.teams, id)
	return nil
}

type InMemoryPlayerRepository struct {
	players map[string]*Player
	mu      sync.RWMutex
}

func NewInMemoryPlayerRepository() PlayerRepository {
	return &InMemoryPlayerRepository{
		players: make(map[string]*Player),
	}
}

func (r *InMemoryPlayerRepository) Save(player *Player) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.players[player.ID] = player
	return nil
}

func (r *InMemoryPlayerRepository) FindByID(id string) (*Player, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	player, ok := r.players[id]
	if !ok {
		return nil, fmt.Errorf("player with ID %s not found", id)
	}
	return player, nil
}

func (r *InMemoryPlayerRepository) FindAll() ([]*Player, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	players := make([]*Player, 0, len(r.players))
	for _, player := range r.players {
		players = append(players, player)
	}
	return players, nil
}

func (r *InMemoryPlayerRepository) Update(player *Player) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.players[player.ID] = player
	return nil
}

func (r *InMemoryPlayerRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.players, id)
	return nil
}
