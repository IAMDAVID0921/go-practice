package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// func main() {
// 	dataChan := make(chan int) // unbuffered channel, there's no space inside
// 	// dataChan := make(chan int, 1)
// 	// expect `fatal error: all goroutines are asleep - deadlock!`
// 	// dataChan <- 789
// 	go func() {
// 		dataChan <- 789
// 	}()
// 	n := <-dataChan
// 	fmt.Printf("n is: %d\n", n)
// }

// 2 ways, we either give space to the channel to make it as a buffered channel, or we
// create a separate goroutine to run `dataChan <- 789`

// if you want to receive data, you have to simultaneously send data to the channel

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func main() {
	dataChan := make(chan int)

	// running on a background thread
	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				// wrapped in another goroutine and use waitgroup to speed up
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan) // to prevent deadlock after 99
	}()

	// running on a main thread
	for n := range dataChan {
		fmt.Printf("n is: %d\n", n)
	}
}
