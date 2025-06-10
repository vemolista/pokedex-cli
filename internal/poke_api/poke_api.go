package poke_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2/"

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationAreas(url *string) (LocationAreasResponse, error) {
	fullUrl := baseUrl + "location-area/"
	if url != nil {
		fullUrl = *url
	}

	i, ok := c.cache.Get(fullUrl)
	if ok {
		var locationAreasResponse LocationAreasResponse
		err := json.Unmarshal(i, &locationAreasResponse)

		if err != nil {
			return LocationAreasResponse{}, fmt.Errorf("error unmarshalling response body: %w", err)
		}

		return locationAreasResponse, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error requesting: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationAreasResponse{}, fmt.Errorf("API returned non-ok status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error reading response body: %w", err)
	}

	c.cache.Add(fullUrl, body)

	var locationAreasResponse LocationAreasResponse
	err = json.Unmarshal(body, &locationAreasResponse)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return locationAreasResponse, nil
}
