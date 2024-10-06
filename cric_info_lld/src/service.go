package src

import (
	"errors"
	"time"
)

type ICricketInfoService interface {
	CreateMatch(homeTeamID string, awayTeamID string, date time.Time, venue string) (*Match, error)
	StartMatch(matchID string) error
	UpdateScore(matchID string, runs int, wickets int) error
	AddCommentary(matchID string, comment string) error
	EndMatch(matchID string) error
	GetMatchDetails(matchID string) (*Match, error)
	GetUpcomingMatches() ([]*Match, error)
	GetCompletedMatches() ([]*Match, error)
	CreateTeam(name string) (*Team, error)
	GetTeamDetails(teamID string) (*Team, error)
	CreatePlayer(name string, teamID string) (*Player, error)
	GetPlayerDetails(playerID string) (*Player, error)
	SearchMatches(query string) ([]*Match, error)
	SearchTeams(query string) ([]*Team, error)
	SearchPlayers(query string) ([]*Player, error)
}

type CricketInfoService struct {
	matchRepo          MatchRepository
	teamRepo           TeamRepository
	playerRepo         PlayerRepository
	idGenerator        IdGenerationStrategy
	scoringStrategy    ScoringStrategy
	commentaryStrategy CommentaryStrategy
}

func NewCricketInfoService(
	matchRepo MatchRepository,
	teamRepo TeamRepository,
	playerRepo PlayerRepository,
	idGenerator IdGenerationStrategy,
	scoringStrategy ScoringStrategy,
	commentaryStrategy CommentaryStrategy,
) ICricketInfoService {
	return &CricketInfoService{
		matchRepo:          matchRepo,
		teamRepo:           teamRepo,
		playerRepo:         playerRepo,
		idGenerator:        idGenerator,
		scoringStrategy:    scoringStrategy,
		commentaryStrategy: commentaryStrategy,
	}
}

func (s *CricketInfoService) CreateMatch(homeTeamID string, awayTeamID string, date time.Time, venue string) (*Match, error) {
	homeTeam, err := s.teamRepo.FindByID(homeTeamID)
	if err != nil {
		return nil, err
	}

	awayTeam, err := s.teamRepo.FindByID(awayTeamID)
	if err != nil {
		return nil, err
	}

	match := NewMatch(homeTeam, awayTeam, date, venue, s.idGenerator.GenerateId())
	err = s.matchRepo.Save(match)
	if err != nil {
		return nil, err
	}

	return match, nil
}

func (s *CricketInfoService) StartMatch(matchID string) error {
	match, err := s.matchRepo.FindByID(matchID)
	if err != nil {
		return err
	}

	match.mu.Lock()
	defer match.mu.Unlock()

	if match.Status != Scheduled {
		return errors.New("match is not in scheduled state")
	}

	match.Status = Live
	return s.matchRepo.Update(match)
}

func (s *CricketInfoService) UpdateScore(matchID string, runs int, wickets int) error {
	match, err := s.matchRepo.FindByID(matchID)
	if err != nil {
		return err
	}

	return s.scoringStrategy.UpdateScore(match, runs, wickets)
}

func (s *CricketInfoService) AddCommentary(matchID string, comment string) error {
	match, err := s.matchRepo.FindByID(matchID)
	if err != nil {
		return err
	}

	return s.commentaryStrategy.AddCommentary(match, comment)
}

func (s *CricketInfoService) EndMatch(matchID string) error {
	match, err := s.matchRepo.FindByID(matchID)
	if err != nil {
		return err
	}

	match.mu.Lock()
	defer match.mu.Unlock()

	if match.Status != Live {
		return errors.New("match is not in live state")
	}

	match.Status = Completed
	return s.matchRepo.Update(match)
}

func (s *CricketInfoService) GetMatchDetails(matchID string) (*Match, error) {
	return s.matchRepo.FindByID(matchID)
}

func (s *CricketInfoService) GetUpcomingMatches() ([]*Match, error) {
	allMatches, err := s.matchRepo.FindAll()
	if err != nil {
		return nil, err
	}

	upcomingMatches := make([]*Match, 0)
	for _, match := range allMatches {
		if match.Status == Scheduled {
			upcomingMatches = append(upcomingMatches, match)
		}
	}

	return upcomingMatches, nil
}

func (s *CricketInfoService) GetCompletedMatches() ([]*Match, error) {
	allMatches, err := s.matchRepo.FindAll()
	if err != nil {
		return nil, err
	}

	completedMatches := make([]*Match, 0)
	for _, match := range allMatches {
		if match.Status == Completed {
			completedMatches = append(completedMatches, match)
		}
	}

	return completedMatches, nil
}

func (s *CricketInfoService) CreateTeam(name string) (*Team, error) {
	team := NewTeam(name, s.idGenerator.GenerateId())
	err := s.teamRepo.Save(team)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (s *CricketInfoService) GetTeamDetails(teamID string) (*Team, error) {
	return s.teamRepo.FindByID(teamID)
}

func (s *CricketInfoService) CreatePlayer(name string, teamID string) (*Player, error) {
	team, err := s.teamRepo.FindByID(teamID)
	if err != nil {
		return nil, err
	}
	player := NewPlayer(name, team, s.idGenerator.GenerateId())
	err = s.playerRepo.Save(player)
	if err != nil {
		return nil, err
	}
	team.Players = append(team.Players, player)
	err = s.teamRepo.Update(team)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (s *CricketInfoService) GetPlayerDetails(playerID string) (*Player, error) {
	return s.playerRepo.FindByID(playerID)
}

func (s *CricketInfoService) SearchMatches(query string) ([]*Match, error) {
	allMatches, err := s.matchRepo.FindAll()
	if err != nil {
		return nil, err
	}
	var matches []*Match
	for _, match := range allMatches {
		if match.HomeTeam.Name == query || match.AwayTeam.Name == query || match.Venue == query {
			matches = append(matches, match)
		}
	}
	return matches, nil
}

func (s *CricketInfoService) SearchTeams(query string) ([]*Team, error) {
	allTeams, err := s.teamRepo.FindAll()
	if err != nil {
		return nil, err
	}
	var teams []*Team
	for _, team := range allTeams {
		if team.Name == query {
			teams = append(teams, team)
		}
	}
	return teams, nil
}

func (s *CricketInfoService) SearchPlayers(query string) ([]*Player, error) {
	allPlayers, err := s.playerRepo.FindAll()
	if err != nil {
		return nil, err
	}
	var players []*Player
	for _, player := range allPlayers {
		if player.Name == query {
			players = append(players, player)
		}
	}
	return players, nil
}
