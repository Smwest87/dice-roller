package roller

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func RollD6(ctx context.Context, diceValueChannel chan int, rollerWaitGroup *sync.WaitGroup) {

	rand.NewSource(time.Now().UnixNano())
	minimum := 1
	maximum := 6
	diceValue := rand.Intn(maximum-minimum+1) + minimum
	if diceValue < 1 {
		errorDetail := fmt.Sprintf("diceValue cannot equal 0\n"+
			"diceValue: %v", diceValue)
		err := errors.New(errorDetail)
		log.Printf("%v", err)
	} else if diceValue > 6 {
		errorDetail := fmt.Sprintf("diceValue cannot be greater than 6\n"+
			"diceValue: %v", diceValue)
		err := errors.New(errorDetail)
		log.Printf("%v", err)
	}
	diceValueChannel <- diceValue
	rollerWaitGroup.Done()
}

func RollAllDice(ctx context.Context, diceTotal int, rollerWaitGroup *sync.WaitGroup, diceValueChannel chan int) {
	for diceCountIterator := 0; diceCountIterator < diceTotal; diceCountIterator++ {
		if len(diceValueChannel) == cap(diceValueChannel) {
			log.Printf("channel is full! Restarting loop.\n Channel Length: %v\nChannel Capacity: %v\n", len(diceValueChannel), cap(diceValueChannel))
			diceCountIterator -= 1
		} else {
			rollerWaitGroup.Add(1)
			go RollD6(ctx, diceValueChannel, rollerWaitGroup)
		}
	}
}
