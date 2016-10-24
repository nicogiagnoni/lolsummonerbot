package services

import (
	"encoding/json"
	"fmt"

	"github.com/nicogiagnoni/lolbot/riot"
	"github.com/nicogiagnoni/lolbot/riot/domain"
	"gopkg.in/resty.v0"
)

//api/lol/{region}/v1.4/summoner/by-name/{summonerNames}
func GetSummoner(summonerName string, region string) (*domain.Summoner, error) {
	url := "https://na.api.pvp.net/api/lol/" + region + "/v1.4/summoner/by-name/" + summonerName + "?api_key=" + riot.GetAPIKey()
	resp, er := resty.R().Get(url)

	if er != nil {
		return nil, er
	} else if resp != nil && resp.StatusCode() < 200 || resp.StatusCode() > 299 {
		switch resp.StatusCode() {
		case 404:
			return nil, fmt.Errorf("Summoner not found")
		default:
			return nil, fmt.Errorf("Error obtaining summoner info")
		}
	}

	var objmap map[string]*json.RawMessage
	json.Unmarshal(resp.Body(), &objmap)

	var s *domain.Summoner
	err := json.Unmarshal(*objmap[summonerName], &s)

	if err != nil {
		return nil, err
	}

	return s, nil
}
