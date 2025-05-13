package commands

import (
	"errors"
	"fmt"
	"math/rand"
)

func CommandCatch(config *PokedexConfig, arguments []string) error {
	if len(arguments) == 0 {
		fmt.Println("Usage: catch <Pokemonname>")
		return errors.New("pokemon name not provided")
	}
	pokemonName := arguments[0]
	poke, err := getPokemonDetails(pokemonName, config.Cache)
	if err != nil {
		fmt.Println("Pokemon " + pokemonName + " not found")
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if err != nil {
		return err
	}
	randomChance := rand.Intn(500)
	threshold := 500 - poke.BaseExperience
	if randomChance >= threshold {
		fmt.Printf("%s was caught\n", pokemonName)
		config.Pokemons[pokemonName] = poke
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s was escaped\n", pokemonName)
	}
	return nil
}
