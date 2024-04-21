package main

type appState string

const (
	MENU appState = "MENU"
	GAME appState = "GAME"
)

type gameState string

const (
	EXPLORE   gameState = "EXPLORE"
	COMBAT    gameState = "COMBAT"
	SHOP      gameState = "SHOP"
	INVENTORY gameState = "INVENTORY"
)

type charSelectState string

const (
	KNIGHT charSelectState = "1"
	MAGE   charSelectState = "2"
	ARCHER charSelectState = "3"
)

type state interface {
	getCommands() map[string]cliCommand
}

type appStateStruct struct {
	state appState
}

type gameStateStruct struct {
	state gameState
}

type charSelectStateStruct struct {
	state charSelectState
}

func initializeMenu() *appStateStruct {
	return &appStateStruct{
		state: MENU,
	}
}

func (as *appStateStruct) getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "get info on actions available",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit game",
			callback:    commandExit,
		},
		"play": {
			name:        "play",
			description: "start game",
			callback:    commandPlay,
		},
	}
}

func (gs *gameStateStruct) getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "get info on actions available",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit game",
			callback:    commandExit,
		},
		"roll": {
			name:        "roll",
			description: "given a dice, roll it",
			callback:    commandRoll,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect your character",
			callback:    commandInspect,
		},
	}
}

func (css *charSelectState) getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "get info on characters available",
			callback:    commandHelp,
		},
		"select": {
			name:        "select",
			description: "select character to play",
			callback:    commandExit,
		},
		"exit": {
			name:        "exit",
			description: "exit game",
			callback:    commandExit,
		},
	}
}
