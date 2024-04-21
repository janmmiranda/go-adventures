package main

import (
	"fmt"

	"github.com/peterh/liner"
)

func main() {
	l := liner.NewLiner()
	defer func() {
		l.Close()
		if err := recover(); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}()

	cfg := &config{
		line:         l,
		currentState: initializeMenu(),
	}

	startRepl(cfg)
}
