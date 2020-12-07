package dice

import (
	"testing"
)

func TestParseDice(t *testing.T) {
	var d Dice = "3d6"
	p := d.Parse()

	if p.NumDice != 3 {
		t.Error("Expected 3 dice, got ", p.NumDice)
	}

	if p.DiceSize != 6 {
		t.Error("Expected 6-sided die, got ", p.DiceSize)
	}

	var anotherDie Dice = "2d10"
	p = anotherDie.Parse()

	if p.NumDice != 2 {
		t.Error("Expected 2 dice, got ", p.NumDice)
	}

	if p.DiceSize != 10 {
		t.Error("Expected 10-sided die, got ", p.DiceSize)
	}
}

func TestRoll(t *testing.T) {
	SeedRandom()

	var d Dice = "3d6"
	var res int
	var rollData []int

	for i := 0; i < 1000; i++ {
		res = d.Roll()
		t.Logf("Roll result: %d\n", res)

		if res > 18 || res < 3 {
			t.Error("Rolled outside range of 3-18: ", res)
		}

		rollData = append(rollData, res)
	}

	var histDict = HistGen(rollData)
	t.Logf("Dict: %d\n", histDict)
}

func TestAnotherRoll(t *testing.T) {
	SeedRandom()

	var d Dice = "2d4"
	var res int
	var rollData []int

	for i := 0; i < 1000; i++ {
		res = d.Roll()
		t.Logf("Rolled 2d4: %d\n", res)

		if res > 8 || res < 2 {
			t.Error("Rolled outside of range 2-8: ", res)
		}

		rollData = append(rollData, res)
	}

	var histDict = HistGen(rollData)
	t.Logf("Dict: %d\n", histDict)
}

func TestBigFatMegaCrunchyRoll(t *testing.T) {
	SeedRandom()

	var d Dice = "8d20"
	var res int
	var rollData []int

	for i := 0; i < 1000; i++ {
		res = d.Roll()
		t.Logf("Rolled 8d20: %d\n", res)

		if res > 160 || res < 8 {
			t.Error("Rolled outside of range 8-160: ", res)
		}

		rollData = append(rollData, res)
	}

	var histDict = HistGen(rollData)
	t.Logf("Dict: %d\n", histDict)
}

func TestRollWithModifier(t *testing.T) {
	SeedRandom()

	var d Dice = "1d20"
	var res int

	for i := 0; i < 100; i++ {
		res = d.RollWithModifier("+2")
		t.Logf("%d\n", res)

		if res > 22 || res < 3 {
			t.Error("Rolled outside range of 3-22: ", res)
		}
	}
}
