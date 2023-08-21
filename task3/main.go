package main

import (
	"fmt"
	"sync"
)

func produce(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i // Sending the number to the channel
	}
	close(ch)
}

func consume(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("%d\n", num)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)

	go produce(ch, &wg)
	go consume(ch, &wg)

	wg.Wait()

	fmt.Println("\nDone")
}
