package main

import (
	"fmt"
	"github.com/smwest87/dice-roller/internal/roller"
	"sync"
)

func main() {
	diceTotal := 10
	diceRollerChannel := make(chan int, diceTotal)
	rollerWaitGroup := sync.WaitGroup{}
	roller.RollAllDice(diceTotal, &rollerWaitGroup, diceRollerChannel)
	rollerWaitGroup.Wait()
	if len(diceRollerChannel) < 10 {
		fmt.Printf("fewer values in in channel than expected: %v", len(diceRollerChannel))
	}
	diceValues := []int{}
	for i := 0; i <= len(diceRollerChannel); i++ {
		diceValues = append(diceValues, <-diceRollerChannel)
	}

	for _, value := range diceValues {
		fmt.Printf("Dice Rolled: %v\n", value)
	}

}
