package services

import (
	"encoding/json"
	"fmt"

	"github.com/nicogiagnoni/lolbot/riot"
	"github.com/nicogiagnoni/lolbot/riot/domain"
	"gopkg.in/resty.v0"
)

func GetLeagueBySummoner(summonerId string, region string) (*domain.League, error) {
	///api/lol/{region}/v1.4/summoner/by-name/{summonerNames}
	url := "https://las.api.pvp.net/api/lol/" + region + "/v2.5/league/by-summoner/" + summonerId + "/entry?api_key=" + riot.GetAPIKey()
	resp, er := resty.R().Get(url)

	if er != nil {
		return nil, er
	} else if resp != nil && resp.StatusCode() < 200 || resp.StatusCode() > 299 {
		switch resp.StatusCode() {
		case 404:
			return nil, fmt.Errorf("Summoner ranked info not found")
		default:
			return nil, fmt.Errorf("Error obtaining summoner ranked info")
		}
	}

	var objmap map[string]*json.RawMessage
	json.Unmarshal(resp.Body(), &objmap)
	var ls []domain.League
	err := json.Unmarshal(*objmap[summonerId], &ls)

	if err != nil {
		return nil, err
	}

	return &ls[0], nil
}
