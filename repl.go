package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	// cfg.Pokedex = make(map[string]apicaller.PokemonInfo)

	for {
		prompt := "Pokedex"
		if cfg.Trainer != "" {
			prompt = cfg.Trainer + "'s Pokedex"
		}
		fmt.Printf("%s > ", prompt)
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())

		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]
		var commandArgument string
		if len(cleanedInput) > 1 {
			commandArgument = cleanedInput[1]
		}

		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(cfg, commandArgument)

			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	return strings.Fields(lowered)
}
