package commands

import "Pokedex/internal/pokecache"

type cliCommand struct {
	name        string
	description string
	Callback    func(*PokedexConfig, []string) error
}

type Pokemon struct {
	name string
}

type PokedexConfig struct {
	Next     string
	Previous string
	Cache    *pokecache.Cache
	Pokemons map[string]PokemonDetails
}

var Commands = map[string]cliCommand{}

func init() {
	Commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		Callback:    CommandHelp,
	}
	Commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		Callback:    CommandExit,
	}
	Commands["map"] = cliCommand{
		name:        "map",
		description: "Maps Pokemon",
		Callback:    CommandMap,
	}
	Commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Maps Pokemon Back",
		Callback:    CommandMapBack,
	}
	Commands["explore"] = cliCommand{
		name:        "explore",
		description: "Explore Pokemon",
		Callback:    CommandExplore,
	}
	Commands["catch"] = cliCommand{
		name:        "catch",
		description: "Catch Pokemon",
		Callback:    CommandCatch,
	}
	Commands["inspect"] = cliCommand{
		name:        "inspect",
		description: "Inspect Pokemon",
		Callback:    CommandInspect,
	}
	Commands["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "list caught pokemons",
		Callback:    CommandPokedex,
	}
}
