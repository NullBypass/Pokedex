package commands

import "fmt"

func CommandPokedex(config *PokedexConfig, arguments []string) error {
	fmt.Println("Your Pokedex:")
	for pokemonName, _ := range config.Pokemons {
		fmt.Println(" -", pokemonName)
	}
	return nil
}
