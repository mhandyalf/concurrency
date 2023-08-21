package main

import (
	"fmt"
	"sync"
)

func sendNumbers(evenCh, oddCh chan<- int, errCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 25; i++ {
		if i%2 == 0 {
			if i <= 20 {
				evenCh <- i
			} else {
				errCh <- fmt.Errorf("number %d too large than 20", i)
			}
		} else {
			if i <= 20 {
				oddCh <- i
			} else {
				errCh <- fmt.Errorf("number %d too large than 20", i)
			}
		}
	}
	close(evenCh)
	close(oddCh)
	close(errCh)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	evenCh := make(chan int, 10)
	oddCh := make(chan int, 10)
	errCh := make(chan error, 5)

	go sendNumbers(evenCh, oddCh, errCh, &wg)

	wg.Wait()

	for {
		select {
		case even, ok := <-evenCh:
			if !ok {
				evenCh = nil
			} else {
				fmt.Printf("Received even: %d\n", even)
			}
		case odd, ok := <-oddCh:
			if !ok {
				oddCh = nil
			} else {
				fmt.Printf("Received odd: %d\n", odd)
			}
		case err, ok := <-errCh:
			if !ok {
				errCh = nil
			} else {
				fmt.Printf("error: %s\n", err)
			}
		}

		if evenCh == nil && oddCh == nil && errCh == nil {
			break
		}
	}
}
