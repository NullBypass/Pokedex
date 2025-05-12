package commands

import (
	"Pokedex/internal/pokecache"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaDetails struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
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

func exploreLocations(locationIdOrName string, cache *pokecache.Cache) (LocationAreaDetails, error) {
	var locationAreaDetails LocationAreaDetails
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", locationIdOrName)
	bodyBytes, ok := cache.Get(url)
	if !ok {

		res, err := http.Get(url)
		if err != nil {
			return LocationAreaDetails{}, err
		}

		if res.StatusCode > 299 {
			return LocationAreaDetails{}, errors.New("Pokedex returned " + res.Status)
		}
		bodyBytes, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaDetails{}, errors.New("cannot read response body")
		}

		defer res.Body.Close()
		cache.Add(url, bodyBytes)
	}

	err := json.Unmarshal(bodyBytes, &locationAreaDetails)
	if err != nil {
		return LocationAreaDetails{}, errors.New("cannot umarshall response body")
	}

	return locationAreaDetails, nil
}
