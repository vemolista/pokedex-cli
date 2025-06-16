package poke_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ShallowLocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationAreas(url *string) (ShallowLocationAreasResponse, error) {
	fullUrl := BaseUrl + "location-area/"
	if url != nil {
		fullUrl = *url
	}

	i, ok := c.cache.Get(fullUrl)
	if ok {
		var locationAreasResponse ShallowLocationAreasResponse
		err := json.Unmarshal(i, &locationAreasResponse)

		if err != nil {
			return ShallowLocationAreasResponse{}, fmt.Errorf("error unmarshalling response body: %w", err)
		}

		return locationAreasResponse, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return ShallowLocationAreasResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ShallowLocationAreasResponse{}, fmt.Errorf("error requesting: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ShallowLocationAreasResponse{}, fmt.Errorf("API returned non-ok status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ShallowLocationAreasResponse{}, fmt.Errorf("error reading response body: %w", err)
	}

	c.cache.Add(fullUrl, body)

	var locationAreasResponse ShallowLocationAreasResponse
	err = json.Unmarshal(body, &locationAreasResponse)
	if err != nil {
		return ShallowLocationAreasResponse{}, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return locationAreasResponse, nil
}
