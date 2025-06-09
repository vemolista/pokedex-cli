package repl

import (
	"fmt"

	"github.com/vemolista/pokedex-cli/internal/poke_api"
)

func commandMap(config *poke_api.Config) error {
	areas, err := poke_api.GetLocationAreas(config.Next)

	if err != nil {
		return fmt.Errorf("error getting location areas: %w", err)
	}

	config.Previous = areas.Previous
	config.Current = areas.Next

	for _, a := range areas.Results {
		fmt.Println(a.Name)
	}

	return nil
}

func commandMapBack(config *poke_api.Config) error {
	areas, err := poke_api.GetLocationAreas(config.Previous)

	if err != nil {
		return fmt.Errorf("error getting location areas: %w", err)
	}

	config.Previous = areas.Previous
	config.Next = areas.Next

	for _, a := range areas.Results {
		fmt.Println(a.Name)
	}

	return nil
}
