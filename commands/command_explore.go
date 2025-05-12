package commands

import "fmt"

func CommandExplore(config *PokedexConfig, arguments []string) error {
	if len(arguments) == 0 {
		return fmt.Errorf("no location id or name")
	}
	areaDetails, err := exploreLocations(arguments[0], config.Cache)
	if err != nil {
		return err
	}
	for _, encounter := range areaDetails.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
