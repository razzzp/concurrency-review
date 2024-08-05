package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func printLetters() {
	for l := 'a'; l <= 'j'; l++ {
		fmt.Printf("%c\n", l)
	}
}

func task1() {
	go printNumbers()
	go printLetters()
	time.Sleep(time.Millisecond)
}

func task2() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		printNumbers()
		wg.Done()
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		printLetters()
		wg.Done()
	}()

	wg.Wait()
}
