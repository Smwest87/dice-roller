package main

import (
	context2 "context"
	"fmt"
	"github.com/smwest87/dice-roller/internal/roller"
	"sync"
)

func main() {
	ctx := context2.Background()
	maximumGoRoutines := 5
	diceTotal := 132
	diceRollerChannel := make(chan int, maximumGoRoutines)
	rollerWaitGroup := sync.WaitGroup{}
	go roller.RollAllDice(ctx, diceTotal, &rollerWaitGroup, diceRollerChannel)
	diceValues := []int{}
	for readFromChannelIterator := 1; readFromChannelIterator <= diceTotal; readFromChannelIterator++ {
		diceValues = append(diceValues, <-diceRollerChannel)
	}
	rollerWaitGroup.Wait()

	for _, value := range diceValues {
		fmt.Printf("Dice Rolled: %v\n", value)
	}

}
