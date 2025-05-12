package commands

import (
	"Pokedex/internal/pokecache"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type LocationItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAPIResponse struct {
	Count int            `json:"count"`
	Next  string         `json:"next"`
	Prev  string         `json:"previous"`
	Items []LocationItem `json:"results"`
}

func getLocations(locationURL string, cache *pokecache.Cache) (LocationAPIResponse, error) {
	if locationURL == "" {
		return LocationAPIResponse{}, errors.New("no location URL provided")
	}

	var locationResponse LocationAPIResponse
	bodyBytes, ok := cache.Get(locationURL)

	if !ok {
		res, err := http.Get(locationURL)
		if err != nil {
			return LocationAPIResponse{}, err
		}

		bodyBytes, err = io.ReadAll(res.Body)

		if err != nil {
			return LocationAPIResponse{}, errors.New("cannot read response body")
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return LocationAPIResponse{}, errors.New("Pokedex returned " + res.Status)
		}
		cache.Add(locationURL, bodyBytes)
	}

	err := json.Unmarshal(bodyBytes, &locationResponse)
	if err != nil {
		return LocationAPIResponse{}, errors.New("cannot umarshall response body")
	}
	return locationResponse, nil
}
