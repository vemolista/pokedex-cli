package repl

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *Config, pokemonName string) error {
	if pokemonName == "" {
		return errors.New("received an empty string as argument")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	chance := rand.IntN(100) + 1
	if chance < 50 {
		fmt.Printf("%s escaped!\n", pokemonName)

		return nil
	}

	pokemon, err := cfg.PokeClient.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("error getting catch response: %w", err)
	}

	cfg.Pokedex[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)

	return nil
}
