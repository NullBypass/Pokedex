package commands

import (
	"encoding/json"
	"errors"
	"fmt"
)

func CommandInspect(config *PokedexConfig, arguments []string) error {
	if len(arguments) == 0 {
		fmt.Println("pokemon name not provided")
		return errors.New("pokemon name not provided")
	}
	pokemonName := arguments[0]
	pokemonDetailsJSONString, ok := config.Cache.Get(pokemonName)
	if !ok {
		fmt.Println("pokemon not found")
		return errors.New("pokemon not found")
	}
	var pokemonDetails PokemonDetails
	err := json.Unmarshal(pokemonDetailsJSONString, &pokemonDetails)
	if err != nil {
		fmt.Println("Cant marshal pokemon details")
		return err
	}

	fmt.Println("Name:", pokemonName)
	fmt.Println("Height:", pokemonDetails.Height)
	fmt.Println("Weight:", pokemonDetails.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonDetails.Stats {
		fmt.Printf("    -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemonDetails.Types {
		fmt.Printf("    - %s\n", pokeType.Type.Name)
	}
	return nil
}
