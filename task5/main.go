package main

import (
	"fmt"
	"sync"
)

func sendNumbers(evenCh, oddCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 25; i++ {
		if i%2 == 0 && i <= 20 {
			evenCh <- i
		} else if i%2 != 0 && i <= 20 {
			oddCh <- i
		}
	}
	close(evenCh)
	close(oddCh)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	evenCh := make(chan int, 10)
	oddCh := make(chan int, 10)

	go sendNumbers(evenCh, oddCh, &wg)

	wg.Wait()

	for even := range evenCh {
		fmt.Printf("Received even: %d\n", even)
	}

	for odd := range oddCh {
		fmt.Printf("Received odd: %d\n", odd)
	}
}
