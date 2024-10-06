package main

import (
	"cric_info_lld.com/src"
	"fmt"
	"time"
)

func main() {
	idGenerator := src.NewIdGenerationUsingUUID()
	matchRepo := src.NewInMemoryMatchRepository()
	teamRepo := src.NewInMemoryTeamRepository()
	playerRepo := src.NewInMemoryPlayerRepository()
	scoringStrategy := src.NewStandardScoringStrategy()
	commentaryStrategy := src.NewBasicCommentaryStrategy()

	cricketInfoService := src.NewCricketInfoService(
		matchRepo,
		teamRepo,
		playerRepo,
		idGenerator,
		scoringStrategy,
		commentaryStrategy,
	)

	// Create teams
	indiaTeam := src.NewTeam("India", idGenerator.GenerateId())
	australiaTeam := src.NewTeam("Australia", idGenerator.GenerateId())

	teamRepo.Save(indiaTeam)
	teamRepo.Save(australiaTeam)

	// Create a match
	match, err := cricketInfoService.CreateMatch(indiaTeam.ID, australiaTeam.ID, time.Now().Add(24*time.Hour), "Sydney Cricket Ground")
	if err != nil {
		fmt.Printf("Error creating match: %v\n", err)
		return
	}

	fmt.Printf("Created match: %+v\n", match)

	// Start the match
	err = cricketInfoService.StartMatch(match.ID)
	if err != nil {
		fmt.Printf("Error starting match: %v\n", err)
		return
	}

	// Update score
	err = cricketInfoService.UpdateScore(match.ID, 10, 1)
	if err != nil {
		fmt.Printf("Error updating score: %v\n", err)
		return
	}

	// Add commentary
	err = cricketInfoService.AddCommentary(match.ID, "What a fantastic shot!")
	if err != nil {
		fmt.Printf("Error adding commentary: %v\n", err)
		return
	}

	// Get match details
	updatedMatch, err := cricketInfoService.GetMatchDetails(match.ID)
	if err != nil {
		fmt.Printf("Error getting match details: %v\n", err)
		return
	}

	fmt.Printf("Updated match: %+v\n", updatedMatch)
	fmt.Printf("Score: %+v\n", updatedMatch.Score)
	fmt.Printf("Commentary: %v\n", updatedMatch.Commentary)

	// End the match
	err = cricketInfoService.EndMatch(match.ID)
	if err != nil {
		fmt.Printf("Error ending match: %v\n", err)
		return
	}

	fmt.Println("Match ended successfully")
}
