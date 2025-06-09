package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vemolista/pokedex-cli/internal/poke_api"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	var config poke_api.Config
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		command := cleanInput(input)[0]

		i, ok := getCommands()[command]
		if !ok {
			fmt.Print("Unknown command.")
		}

		err := i.callback(&config)
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
	callback    func(config *poke_api.Config) error
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
	}
}
