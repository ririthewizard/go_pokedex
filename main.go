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
			name: "exit",
			description: "Exit the poxedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "	",
		}
	}
	scanner := bufio.NewScanner(os.Stdin)


	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		commandExit()
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words
}

func commandExit() error {
	fmt.Print("Closing the Poxedex... Goodbye!")
	os.Exit(0)
	return nil
}
