package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands = map[string]cliCommand{
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
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		command := cleanInput(input)[0]

		i, ok := commands[command]
		if !ok {
			fmt.Print("Unknown command.")
		}

		i.callback()
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	// for _, v := range commands {
	// 	fmt.Printf("%s: %s", v.name, v.description)
	// }

	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")

	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
