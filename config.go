package main

import (
	"github.com/janmmiranda/go_adventures/internal/entities"
	"github.com/peterh/liner"
)

type config struct {
	line         *liner.State
	currentState state
	player       *entities.Player
}
