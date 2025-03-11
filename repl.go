package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		fields := strings.Fields(strings.ToLower(input))

		userCommand := getCmd(fields[0])
		sysCommand, ok := getCommands()[userCommand]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		var arg string
		if len(fields) > 1 {
			arg = fields[1]
		}

		err := sysCommand.callback(cmdParams{
			cfg:   &cfg,
			cache: &cache,
			arg:   arg,
		})
		if err != nil {
			fmt.Println(err.Error())
		}

		// extra line to add some white space between commands
		fmt.Println()
	}
}
