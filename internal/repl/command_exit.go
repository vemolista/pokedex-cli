package repl

import (
	"fmt"
	"os"

	"github.com/vemolista/pokedex-cli/internal/poke_api"
)

func commandExit(config *poke_api.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}
