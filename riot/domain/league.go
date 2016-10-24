package domain

import (
	"strconv"
)

type League struct {
	Tier    string
	Name    string
	Entries []LeagueEntry
}

type LeagueEntry struct {
	Division     string
	IsFreshBlood bool
	IsHotStreak  bool
	IsInactive   bool
	IsVeteran    bool
	LeaguePoints int64
	Losses       int64
	Wins         int64
	Playstyle    string
	MiniSeries   *MiniSeries
}

type MiniSeries struct {
	Losses   int64
	Wins     int64
	Target   int64
	Progress string
}

func (le *LeagueEntry) GetLeaguePoints() string {
	return strconv.FormatInt(le.LeaguePoints, 10)
}

func (le *LeagueEntry) GetWins() string {
	return strconv.FormatInt(le.Wins, 10)
}

func (le *LeagueEntry) GetLosses() string {
	return strconv.FormatInt(le.Losses, 10)
}
