package repl

import "fmt"

func commandPokedex(cfg *Config, arg string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("You have no pokemons in your Pokedex.")

		return nil
	}

	fmt.Println("Your Pokedex:")
	for k := range cfg.Pokedex {
		fmt.Printf("- %v\n", k)
	}

	return nil
}
