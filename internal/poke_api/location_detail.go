package poke_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationDetails(locationArea string) (LocationAreaResponse, error) {
	url := BaseUrl + "location-area/" + locationArea

	// i, ok := c.cache.Get(url)
	// if ok {
	// 	var locationAreasResponse ShallowLocationAreasResponse
	// 	err := json.Unmarshal(i, &locationAreasResponse)

	// 	if err != nil {
	// 		return ShallowLocationAreasResponse{}, fmt.Errorf("error unmarshalling response body: %w", err)
	// 	}

	// 	return locationAreasResponse, nil
	// }

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error requesting: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationAreaResponse{}, fmt.Errorf("API returned non-ok status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error reading response body: %w", err)
	}

	c.cache.Add(url, body)

	var locationAreaResponse LocationAreaResponse
	err = json.Unmarshal(body, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return locationAreaResponse, nil
}
