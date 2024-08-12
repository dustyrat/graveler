package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

const batchSize = 100

var items = []int{1, 2, 3, 4}

func main() {
	run(1000000000, 231, 177)
}

func run(maxSessions, maxRolls, target int) {
	fmt.Printf("Running %d roll sessions of %d rolls with a target of %d\n", maxSessions, maxRolls, 177)
	start := time.Now()

	sessions := 0
	maxOnes := 0
	defer func() {
		fmt.Println("***** DONE *****")
		fmt.Println("Highest Ones Roll:", maxOnes)
		fmt.Println("Number of Roll Sessions:", sessions)
		fmt.Println("Time taken:", time.Now().Sub(start))
	}()

	ch := make(chan struct{}, batchSize)
	var wg sync.WaitGroup
	var mux sync.Mutex
	for maxOnes < target && sessions < maxSessions {
		wg.Add(1)
		ch <- struct{}{}
		go func() {
			defer wg.Done()
			numbers := []int{0, 0, 0, 0}
			for i := 0; i < maxRolls; i++ {
				roll := items[rand.Intn(len(items))]
				numbers[roll-1]++
			}
			mux.Lock()
			if numbers[0] > maxOnes {
				maxOnes = numbers[0]
				fmt.Printf("New Max: %d Session: %d\n", maxOnes, sessions)
			}
			mux.Unlock()
			<-ch
		}()
		sessions++

		if sessions%(maxSessions/100) == 0 {
			fmt.Printf("%5.1f %% %*d/%d - Max: %d (%s)\n", (float64(sessions)/float64(maxSessions))*100, len(strconv.Itoa(maxSessions)), sessions, maxSessions, maxOnes, time.Now().Sub(start))
		}
	}
	wg.Wait()
}
