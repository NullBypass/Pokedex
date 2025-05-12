package commands

import "fmt"

func CommandMap(config *PokedexConfig) error {
	locationResponse, err := getLocations(config.Next, config.Cache)
	if err != nil {
		return err
	}
	for _, item := range locationResponse.Items {
		fmt.Println(item.Name)
	}
	config.Next = locationResponse.Next
	config.Previous = locationResponse.Prev
	return nil
}

func CommandMapBack(config *PokedexConfig) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	locationResponse, err := getLocations(config.Previous, config.Cache)
	if err != nil {
		return err
	}
	for _, item := range locationResponse.Items {
		fmt.Println(item.Name)
	}
	config.Next = locationResponse.Next
	config.Previous = locationResponse.Prev
	return nil
}
