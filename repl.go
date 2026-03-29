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

func scrollRepl(items []string, index int) (string, error) {
	if index >= len(items) {
		return "", fmt.Errorf("Index %d is out of range of slice", index)
	}

	menuScanner := bufio.NewScanner(os.Stdin)

	for {
		for idx, item := range items {
			margin := "  "
			if idx == index {
				margin = "> "
			}
			fmt.Printf("%s%s\n", margin, item)
		}
		menuScanner.Scan()
		ok, direction := parseArrow(menuScanner.Text())

	}

	return "", nil
}

func parseArrow(text string) (contains, direction bool) {
	if strings.Contains(text, "^[[A") {
		return true, true
	}
	if strings.Contains(text, "^[[B") {
		return true, false
	}
	return false, false
}
