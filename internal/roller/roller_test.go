package roller

import (
	"sync"
	"testing"
)

// TestRollD6 calls RollD6. Returns a random value of 1-6
func TestRollD6(t *testing.T) {
	diceValueChannel := make(chan int)
	rollerWaitGroup := sync.WaitGroup{}
	rollerWaitGroup.Add(1)
	go RollD6(diceValueChannel, &rollerWaitGroup)
	diceValue := <-diceValueChannel
	close(diceValueChannel)
	if diceValue < 1 || diceValue > 6 {
		t.FailNow()
	}

	rollerWaitGroup.Wait()

}

func TestRollAllDice(t *testing.T) {
	diceTotal := 10
	rollerWaitGroup := sync.WaitGroup{}
	allDiceValue, err := RollAllDice(diceTotal, &rollerWaitGroup)
	if err != nil {
		t.Logf("error present when none was expected, %v", err.Error())
		t.FailNow()
	}
	if len(allDiceValue) > 10 {
		t.Logf("total number of dice is greater than 10, %v", len(allDiceValue))
		t.FailNow()
	} else if len(allDiceValue) < 1 {
		t.Logf("total number of dice is less than 10, %v", len(allDiceValue))
		t.FailNow()

	}
}
