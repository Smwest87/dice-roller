package roller

import (
	"context"
	"sync"
	"testing"
)

// TestRollD6 calls RollD6. Returns a random value of 1-6
func TestRollD6(t *testing.T) {
	ctx := context.Background()
	diceValueChannel := make(chan int)
	rollerWaitGroup := sync.WaitGroup{}
	rollerWaitGroup.Add(1)
	go RollD6(ctx, diceValueChannel, &rollerWaitGroup)
	diceValue := <-diceValueChannel
	close(diceValueChannel)
	if diceValue < 1 || diceValue > 6 {
		t.FailNow()
	}
	rollerWaitGroup.Wait()

}

func TestRollAllDice(t *testing.T) {
	ctx := context.Background()
	diceTotal := 10
	rollerWaitGroup := sync.WaitGroup{}
	rollerChannel := make(chan int, 10)
	RollAllDice(ctx, diceTotal, &rollerWaitGroup, rollerChannel)
	rollerWaitGroup.Wait()
	if len(rollerChannel) > 10 {
		t.Logf("total number of dice is greater than 10, %v", len(rollerChannel))
		t.FailNow()
	} else if len(rollerChannel) < 1 {
		t.Logf("total number of dice is less than 10, %v", len(rollerChannel))
		t.FailNow()
	}
}
