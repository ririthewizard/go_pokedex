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

func main() {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the poxedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		commandWords := cleanInput(input)

		if len(commandWords) > 0 {
			cmdName := commandWords[0]

			if cmd, ok := commands[cmdName]; ok {
				cmd.callback()
			} else {
				fmt.Println("Unknown command. Try typing help for all the commands available")
			}
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words
}

func commandExit() error {
	fmt.Println("Closing the Poxedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Prints a help message\nexit: Exit the Pokedex")
	return nil
}
