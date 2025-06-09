package repl

import (
	"fmt"

	"github.com/vemolista/pokedex-cli/internal/poke_api"
)

func commandHelp(config *poke_api.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, v := range getCommands() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}

	return nil
}
