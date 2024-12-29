package roller

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func RollD6(diceValueChannel chan int, rollerWaitGroup *sync.WaitGroup) {
	defer rollerWaitGroup.Done()
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
}

func RollAllDice(diceTotal int, rollerWaitGroup *sync.WaitGroup, diceValueChannel chan int) {
	for i := 0; i < diceTotal; i++ {
		rollerWaitGroup.Add(1)
		go RollD6(diceValueChannel, rollerWaitGroup)
	}
}
