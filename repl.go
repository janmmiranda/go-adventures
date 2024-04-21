package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/janmmiranda/go_adventures/internal/common"
)

func startRepl(cfg *config) {
	if f, err := os.Open(common.ReplHistory); err == nil {
		cfg.line.ReadHistory(f)
		f.Close()
	}

	fmt.Println(common.IntroText)

	for {
		if input, err := cfg.line.Prompt("Go > "); err == nil {
			cfg.line.AppendHistory(input)
			words := cleanInput(input)

			if len(words) == 0 {
				continue
			}

			commandName := words[0]
			args := []string{}
			if len(words) > 1 {
				args = words[1:]
			}

			command, exists := cfg.currentState.getCommands()[commandName]
			if exists {
				err := command.callback(cfg, args...)
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else {
				fmt.Println("Unkown command")
				continue
			}
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
