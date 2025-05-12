package main

import (
	"Pokedex/commands"
	"bufio"
	"fmt"
	"os"
)

const PROMPT = "Pokedex > "

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := commands.PokedexConfig{
		Next: "https://pokeapi.co/api/v2/location-area",
	}
	for {
		fmt.Print(PROMPT)
		scanner.Scan()
		line := scanner.Text()
		cleanedLine := cleanInput(line)
		if len(cleanedLine) == 0 {
			continue
		}

		command := cleanedLine[0]
		c, ok := commands.Commands[command]
		if !ok {
			fmt.Printf("Unknown command: %s\n", command)
		} else {
			err := c.Callback(&config)
			if err != nil {
				return
			}
		}
	}
}
