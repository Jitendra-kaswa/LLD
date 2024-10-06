package src

type MatchStatus string

const (
	Scheduled MatchStatus = "Scheduled"
	Live      MatchStatus = "Live"
	Completed MatchStatus = "Completed"
)
