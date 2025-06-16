package repl

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *Config, pokemonName string) error {
	if pokemonName == "" {
		return errors.New("received an empty string as argument")
	}

	pokemon, ok := cfg.Pokedex[pokemonName]
	if !ok {
		fmt.Printf("you have not caught %s!\n", pokemonName)

		return nil
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, i := range pokemon.Stats {
		fmt.Printf("- %v: %v\n", i.Stat.Name, i.BaseStat)
	}

	fmt.Println("Types:")
	for _, i := range pokemon.Types {
		fmt.Printf("- %v\n", i.Type.Name)
	}

	return nil
}
