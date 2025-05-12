package commands

type cliCommand struct {
	name        string
	description string
	Callback    func(config *PokedexConfig) error
}

type PokedexConfig struct {
	Next     string
	Previous string
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
}
