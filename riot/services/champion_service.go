package services

import (
	"encoding/json"
	"fmt"

	"github.com/nicogiagnoni/lolbot/riot"
	"github.com/nicogiagnoni/lolbot/riot/domain"
	"gopkg.in/resty.v0"
)

///api/lol/static-data/{region}/v1.2/champion/{id}
func GetChampion(region string, championId string) (*domain.ChampionStatic, error) {
	url := "https://las.api.pvp.net/api/lol/static-data/" + region + "/v1.2/champion/" + championId + "?api_key=" + riot.GetAPIKey()
	resp, er := resty.R().Get(url)

	if er != nil {
		return nil, er
	} else if resp != nil && resp.StatusCode() < 200 || resp.StatusCode() > 299 {
		switch resp.StatusCode() {
		case 404:
			return nil, fmt.Errorf("Champion not found")
		default:
			return nil, fmt.Errorf("Error obtaining champion info")
		}
	}

	var result *domain.ChampionStatic
	err := json.Unmarshal(resp.Body(), &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
