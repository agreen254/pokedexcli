package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/agreen254/pokedexcli/internal/pokecache"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(time.Second * 10)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			panic(err)
		}

		input := scanner.Text()

		userCommand := getCmd(input)
		sysCommand, ok := getCommands()[userCommand]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := sysCommand.callback(&cfg, &cache)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println()
	}
}
