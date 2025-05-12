package commands

import "fmt"

func CommandHelp(config *PokedexConfig) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for name, cmd := range Commands {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}

	return nil
}
