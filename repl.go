package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	fields := strings.Fields(text)

	lowered := make([]string, len(fields))
	for i, field := range fields {
		lowered[i] = strings.ToLower(field)
	}
	return lowered
}

func getCmd(input string) string {
	fields := strings.Fields(strings.ToLower(input))
	if len(fields) == 0 {
		return ""
	} else {
		return fields[0]
	}
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			panic(err)
		}

		input := scanner.Text()
		fmt.Printf("Your command was: %s\n", getCmd(input))
	}
}
