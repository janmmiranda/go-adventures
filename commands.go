package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/janmmiranda/go_adventures/internal/common"
	"github.com/janmmiranda/go_adventures/internal/dices"
	"github.com/janmmiranda/go_adventures/internal/entities"
)

func commandHelp(cfg *config, args ...string) error {
	for _, cmd := range cfg.currentState.getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(cfg *config, args ...string) error {
	if f, err := os.Create(".repl_history"); err == nil {
		cfg.line.WriteHistory(f)
		cfg.line.Close()
		f.Close()
	}

	fmt.Println()
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}

func commandPlay(cfg *config, args ...string) error {
	cfg.currentState = &appStateStruct{
		state: appState(GAME),
	}
	// do character select etc
	commandCharacterSelect(cfg, args...)
	return nil
}

func commandCharacterSelect(cfg *config, args ...string) error {
	fmt.Println("Create your character")
	fmt.Println("What is your name, adventurer?")
	var name string
	if temp, err := cfg.line.Prompt("Go > "); err != nil {
		return errors.New("Error reading name:\n" + err.Error())
	} else {
		name = temp
	}
	fmt.Println("Choose your class!")
	time.Sleep(1 * time.Second)
	fmt.Println(common.KnightSprite)
	fmt.Println("Knight [1]")
	time.Sleep(1 * time.Second)
	fmt.Println(common.MageSprite)
	fmt.Println("Mage [2]")
	time.Sleep(1 * time.Second)
	fmt.Println(common.ArcherSprite)
	fmt.Println("Archer [3]")
	time.Sleep(1 * time.Second)
	fmt.Println("Knight [1] | Mage [2] | Archer [3]")
	var class string
	//make this a seperate function and check to make sure input is either 1, 2, or 3
	if temp, err := cfg.line.Prompt("Go > "); err != nil {
		return errors.New("Error reading class:\n" + err.Error())
	} else {
		class = temp
	}

	if temp, err := entities.CreatePlayer(name, class); err != nil {
		return errors.New("Error creating player:\n" + err.Error())
	} else {
		cfg.player = temp
	}

	cfg.currentState = &gameStateStruct{
		state: gameState(EXPLORE),
	}
	fmt.Printf("Character created\nYou are %s the %s\n", cfg.player.Name, cfg.player.Class)
	return nil
}

func commandRoll(cfg *config, args ...string) error {
	for _, arg := range args {
		tempDice, count := dices.GetDice(arg)
		if tempDice != dices.D0 {
			fmt.Printf("Rolling D%v %v times...\n", tempDice, count)
			if count != 1 {
				fmt.Printf("Rolled %v\n", dices.Roll(0, count, tempDice))
			} else {
				fmt.Printf("Rolled %v\n", tempDice.Roll())
			}
		}
	}
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	// TODO print out Players stats
	if cfg.player == nil {
		return errors.New("Player not set")
	}
	fmt.Printf("%s the %s\n", cfg.player.Name, cfg.player.Class)
	fmt.Printf("Current health: %v\n", cfg.player.CurrentHealth)
	fmt.Printf("Armor Class: %v\n", cfg.player.ArmorClass)
	fmt.Printf("Speed: %vft\n", cfg.player.Speed)
	fmt.Println("Abilities")
	fmt.Printf("- Strength: %v\n", cfg.player.Strength)
	fmt.Printf("- Dexterity: %v\n", cfg.player.Dexterity)
	fmt.Printf("- Constitution: %v\n", cfg.player.Constitution)
	fmt.Printf("- Intelligence: %v\n", cfg.player.Intelligence)
	fmt.Printf("- Wisdom: %v\n", cfg.player.Wisdom)
	fmt.Printf("- Charisma: %v\n", cfg.player.Charisma)
	return nil
}
