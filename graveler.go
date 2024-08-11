package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const batchSize = 100

var items = []int{1, 2, 3, 4}

func main() {
	run(177, 1000000000)
}

func run(target, maxRolls int) {
	fmt.Printf("Running %d rolls...\n", maxRolls)
	start := time.Now()

	rolls := 0
	maxOnes := 0
	defer func() {
		fmt.Println("***** DONE *****")
		fmt.Println("Highest Ones Roll:", maxOnes)
		fmt.Println("Number of Roll Sessions:", rolls)
		fmt.Println("Time taken:", time.Now().Sub(start))
	}()

	ch := make(chan struct{}, batchSize)
	var wg sync.WaitGroup

	for maxOnes < target && rolls < maxRolls {
		wg.Add(1)
		ch <- struct{}{}
		go func() {
			defer wg.Done()
			numbers := []int{0, 0, 0, 0}
			for i := 0; i < 231; i++ {
				roll := items[rand.Intn(len(items))]
				numbers[roll-1]++
			}
			if numbers[0] > maxOnes {
				maxOnes = numbers[0]
			}
			<-ch
		}()
		rolls++

		if rolls%(maxRolls/100) == 0 {
			fmt.Printf("(%s) %d/%d - %0.2f%%\n", time.Now().Sub(start), rolls, maxRolls, (float64(rolls)/float64(maxRolls))*100)
		}
	}
	wg.Wait()
}
