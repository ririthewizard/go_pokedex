package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	cleanInput("hello world")
}

func cleanInput(text string) []string {
	words := make([]string, len(text))
	if len(words) == 0 {
		return []string{}
	}

	for i := 0; i < len(text); i++ {
		words = append(words, words[i])
	}

	fmt.Println(words)

	return []string{}
}
