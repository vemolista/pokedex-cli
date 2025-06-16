package main

import (
	"time"

	"github.com/vemolista/pokedex-cli/internal/poke_api"
	"github.com/vemolista/pokedex-cli/internal/repl"
)

func main() {
	pokeClient := poke_api.NewClient(5*time.Second, 5*time.Minute)

	cfg := &repl.Config{
		PokeClient: pokeClient,
		Pokedex:    make(map[string]poke_api.PokemonResponse),
	}

	repl.StartRepl(cfg)
}
