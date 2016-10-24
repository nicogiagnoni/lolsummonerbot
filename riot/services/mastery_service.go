package services

import (
	"encoding/json"
	"fmt"

	"github.com/nicogiagnoni/lolbot/riot"
	"github.com/nicogiagnoni/lolbot/riot/domain"
	"gopkg.in/resty.v0"
)

//championmastery/location/{location}/player/{playerId}/topchampions
func GetMasteryTopChampions(summonerId string, region string) ([]domain.ChampionMastery, error) {
	url := "https://na.api.pvp.net/championmastery/location/" + region + "/player/" + summonerId + "/topchampions?count=5&api_key=" + riot.GetAPIKey()
	resp, er := resty.R().Get(url)

	if er != nil {
		return nil, er
	} else if resp != nil && resp.StatusCode() < 200 || resp.StatusCode() > 299 {
		switch resp.StatusCode() {
		case 404:
			return nil, fmt.Errorf("Summoner Mastery champs not found")
		default:
			return nil, fmt.Errorf("Error obtaining summoner mastery champs info")
		}
	}

	var result []domain.ChampionMastery
	err := json.Unmarshal(resp.Body(), &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
