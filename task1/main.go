package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for ch := 'a'; ch <= 'j'; ch++ {
		fmt.Printf("%c\n", ch)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printNumbers()
	go printLetters()

	// Sleep for a while to allow goroutines to finish execution
	time.Sleep(2 * time.Second)
	fmt.Println("\nDone")
}
