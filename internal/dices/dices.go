package dices

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

type Dice int

const (
	D0   Dice = 0
	D2   Dice = 2
	D4   Dice = 4
	D6   Dice = 6
	D8   Dice = 8
	D10  Dice = 10
	D12  Dice = 12
	D20  Dice = 20
	D100 Dice = 100
)

func (d Dice) Roll() int {
	return 1 + rand.Intn(int(d))
}

func Roll(modifier int, count int, d Dice) int {
	res := modifier
	for i := 0; i < count; i++ {
		res += d.Roll()
	}
	return res
}

func RollAdvantage(modifier int, d Dice) int {
	return int(math.Max(float64(d.Roll()), float64(d.Roll()))) + modifier
}

func RollDisadvantage(modifier int, d Dice) int {
	return int(math.Min(float64(d.Roll()), float64(d.Roll()))) + modifier
}

func RollPercentile() int {
	roll1 := D10.Roll()
	roll2 := D10.Roll()
	return roll1*10 + roll2
}

func GetDice(input string) (Dice, int) {
	fmt.Println("getting dice")
	count, diceName := splitDiceInput(input)
	switch diceName {
	case "d2":
		return D2, count
	case "d4":
		return D4, count
	case "d6":
		return D6, count
	case "d8":
		return D8, count
	case "d10":
		return D10, count
	case "d12":
		return D12, count
	case "d20":
		return D20, count
	default:
		return D0, count
	}
}

func splitDiceInput(diceInput string) (int, string) {
	parts := strings.SplitN(diceInput, "d", 2)
	if len(parts) != 2 {
		fmt.Println("Invalid input format")
		return 0, ""
	}
	var num int
	if parts[0] == "" {
		num = 1
	} else {
		temp, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Failed to convert number:", err)
			return 0, ""
		} else {
			num = temp
		}
	}
	return num, "d" + parts[1]
}
