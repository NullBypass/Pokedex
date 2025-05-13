package commands

import (
	"errors"
	"fmt"
)

func CommandExplore(config *PokedexConfig, arguments []string) error {
	if len(arguments) == 0 {
		fmt.Println("no location id or name")
		return errors.New("no location id or name")
	}
	areaDetails, err := exploreLocations(arguments[0], config.Cache)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", arguments[0])
	for _, encounter := range areaDetails.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}
