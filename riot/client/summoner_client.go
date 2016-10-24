package client

import "github.com/nicogiagnoni/lolbot/riot/services"

func GetSummonerRanking(summonerName string, region string) (string, error) {
	s, e := services.GetSummoner(summonerName, region)
	if e != nil {
		return "", e
	}

	l, er := services.GetLeagueBySummoner(s.GetId(), region)
	if er != nil {
		return "", er
	}

	tier := l.Tier + " " + l.Entries[0].Division
	ranking := s.GetName()
	ranking += "\n"
	ranking += tier
	ranking += "\n"
	ranking += l.Entries[0].GetLeaguePoints()
	ranking += " LP"
	ranking += " / "
	ranking += l.Entries[0].GetWins()
	ranking += "W "
	ranking += l.Entries[0].GetLosses()
	ranking += "L"

	if l.Entries[0].MiniSeries != nil {
		ranking += "\nPromotion: "
		ranking += l.Entries[0].MiniSeries.Progress
	}

	return ranking, nil
}

func GetSummonerMasteryChamps(summonerName string, region string) (string, error) {
	//Retrieve summoner for ID
	s, e := services.GetSummoner(summonerName, region)
	if e != nil {
		return "", e
	}

	//Retrieve summoner top mastery champions
	l, er := services.GetMasteryTopChampions(s.GetId(), region)
	if er != nil {
		return "", er
	}

	masteryChamps := "Best champions: "

	//Range champs slice and retrieve Champ static data
	for _, v := range l {
		champ, err := services.GetChampion(region, v.GetChampionId())

		if err != nil {
			return "", nil
		}

		masteryChamps += "\n"
		masteryChamps += "- "
		masteryChamps += champ.Name
		masteryChamps += " / Mastery Level "
		masteryChamps += v.GetChampionLevel()
		masteryChamps += " / "
		masteryChamps += v.GetChampionPoints()
		masteryChamps += " Mastery Points"
	}

	return masteryChamps, nil
}
