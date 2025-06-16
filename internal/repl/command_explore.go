package repl

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *Config, area string) error {
	if area == "" {
		return errors.New("received an empty string")
	}

	location, err := cfg.PokeClient.GetLocationDetails(area)
	if err != nil {
		return fmt.Errorf("error getting location details: %w", err)
	}

	fmt.Printf("Exploring %s...\n", area)
	fmt.Printf("Found some Pokemon!\n")
	for _, p := range location.PokemonEncounters {
		fmt.Printf("- %s\n", p.Pokemon.Name)
	}

	return nil
}
