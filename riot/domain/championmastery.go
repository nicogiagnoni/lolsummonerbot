package domain

import "strconv"

type ChampionMastery struct {
	ChampionId                   int64
	ChampionLevel                int64
	ChampionPoints               int64
	ChampionPointsSinceLastLevel int64
	ChampionPointsUntilNextLevel int64
	ChestGranted                 bool
	LastPlayTime                 int64
	PlayerId                     int64
}

func (cm *ChampionMastery) GetChampionId() string {
	return strconv.FormatInt(cm.ChampionId, 10)
}

func (cm *ChampionMastery) GetChampionLevel() string {
	return strconv.FormatInt(cm.ChampionLevel, 10)
}

func (cm *ChampionMastery) GetChampionPoints() string {
	return strconv.FormatInt(cm.ChampionPoints, 10)
}
