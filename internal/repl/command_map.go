package repl

import (
	"fmt"
)

func commandMap(cfg *Config, arg string) error {
	areas, err := cfg.PokeClient.GetLocationAreas(cfg.nextLocationsUrl)

	if err != nil {
		return fmt.Errorf("error getting location areas: %w", err)
	}

	cfg.previousLocationsUrl = areas.Previous
	cfg.nextLocationsUrl = areas.Next

	for _, a := range areas.Results {
		fmt.Println(a.Name)
	}

	return nil
}

func commandMapBack(cfg *Config, arg string) error {
	areas, err := cfg.PokeClient.GetLocationAreas(cfg.previousLocationsUrl)

	if err != nil {
		return fmt.Errorf("error getting location areas: %w", err)
	}

	cfg.previousLocationsUrl = areas.Previous
	cfg.nextLocationsUrl = areas.Next

	for _, a := range areas.Results {
		fmt.Println(a.Name)
	}

	return nil
}
