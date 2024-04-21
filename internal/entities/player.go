package entities

import (
	"errors"
	"fmt"

	"github.com/janmmiranda/go_adventures/internal/dices"
)

type Player struct {
	Name          string
	Class         string
	Level         int
	Inventory     []Item
	CurrentHealth int
	MaxHealth     int
	ArmorClass    int
	Speed         int
	Strength      int
	Dexterity     int
	Constitution  int
	Intelligence  int
	Wisdom        int
	Charisma      int
}

func CreatePlayer(name string, classType string) (*Player, error) {
	switch classType {
	case "1":
		return createKnight(name), nil
	case "2":
		return createMage(name), nil
	case "3":
		return createArcher(name), nil
	default:
		return &Player{}, errors.New("unable to create new character")
	}
}

func createKnight(name string) *Player {
	return &Player{
		Name:          name,
		Class:         "Knight",
		Level:         1,
		Inventory:     []Item{},
		CurrentHealth: 7,
		MaxHealth:     7,
		ArmorClass:    18,
		Speed:         30,
		Strength:      16,
		Dexterity:     9,
		Constitution:  15,
		Intelligence:  11,
		Wisdom:        13,
		Charisma:      14,
	}
}

func createMage(name string) *Player {
	return &Player{
		Name:          name,
		Class:         "Mage",
		Inventory:     []Item{},
		CurrentHealth: 5,
		MaxHealth:     5,
		ArmorClass:    12,
		Speed:         30,
		Strength:      10,
		Dexterity:     15,
		Constitution:  14,
		Intelligence:  16,
		Wisdom:        12,
		Charisma:      8,
	}
}

func createArcher(name string) *Player {
	return &Player{
		Name:          name,
		Class:         "Archer",
		Inventory:     []Item{},
		CurrentHealth: 5,
		MaxHealth:     5,
		ArmorClass:    14,
		Speed:         25,
		Strength:      8,
		Dexterity:     16,
		Constitution:  12,
		Intelligence:  13,
		Wisdom:        10,
		Charisma:      16,
	}
}

func (p *Player) getInventory() []Item {
	return p.Inventory
}

func (p *Player) getModifier(ability int) int {
	return (ability - 10) / 2
}

func (p *Player) getInitiative() int {
	mod := p.getModifier(p.Dexterity)
	return dices.Roll(mod, 1, dices.D20)
}

func (p *Player) talk() {
	fmt.Println("I have nothing to say...")
}

func (p *Player) getHealth() int {
	return p.CurrentHealth
}
