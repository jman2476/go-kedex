package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var registry = map[string]cliCommand{
	// "exit": {
	// 	name:        "exit",
	// 	description: "Exit the Pokedex",
	// 	callback:    commandExit,
	// },
	// "help": {
	// 	name:        "help",
	// 	description: "Displays a help message",
	// 	callback:    commandHelp,
	// },
}

func init() {
	registry["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	registry["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		command, ok := registry[cleaned[0]]
		if ok {
			command.callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	return strings.Fields(lowered)
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")

	for _, command := range registry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
