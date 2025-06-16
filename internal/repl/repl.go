package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vemolista/pokedex-cli/internal/poke_api"
)

type Config struct {
	PokeClient           poke_api.Client
	nextLocationsUrl     *string
	previousLocationsUrl *string
	Pokedex              map[string]poke_api.PokemonResponse
}

func StartRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		tokens := cleanInput(input)
		command := tokens[0]

		arg := ""
		if len(tokens) > 1 {
			arg = tokens[1]
		}

		i, ok := getCommands()[command]
		if !ok {
			fmt.Printf("Unknown command '%s'.\n", input)
			fmt.Println("Type 'help' for a list of available commands.")

			continue
		}

		err := i.callback(cfg, arg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display help message.",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get next locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous locations.",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location in detail.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon.",
			callback:    commandInspect,
		},
	}
}
