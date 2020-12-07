package dice

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type (
	Dice       string
	ParsedDice struct {
		NumDice  int
		DiceSize int
	}
)

func (dice *Dice) Parse() ParsedDice {
	arr := strings.Split(string(*dice), "d")
	numDice, _ := strconv.Atoi(arr[0])
	diceSize, _ := strconv.Atoi(arr[1])

	return ParsedDice{numDice, diceSize}
}

func (dice *Dice) Roll() (roll int) {
	parsed := dice.Parse()
	roll = 0
	for i := 1; i <= parsed.NumDice; i++ {
		min := 1
		max := parsed.DiceSize + 1
		roll += dice.Random(min, max)
	}
	return roll
}

func (dice *Dice) RollWithModifier(mod string) int {
	modifier, _ := strconv.Atoi(mod)
	return dice.Roll() + modifier
}

func (dice *Dice) Random(min, max int) int {
	return rand.Intn(max-min) + min
}

func SeedRandom() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func HistGen(data []int) map[int]int {
	var histData = make(map[int]int)
	for _, number := range data {
		histData[number] = histData[number] + 1
	}
	return histData
}
