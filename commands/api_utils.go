package commands

import (
	"encoding/json"
	"errors"
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

func getLocations(locationURL string) (LocationAPIResponse, error) {
	res, err := http.Get(locationURL)
	if err != nil {
		return LocationAPIResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationAPIResponse{}, errors.New("Pokedex returned " + res.Status)
	}

	var locationResponse LocationAPIResponse

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationResponse)
	if err != nil {
		return LocationAPIResponse{}, err
	}

	return locationResponse, nil
}
