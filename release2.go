package main

import (
	"fmt"
	"sync"
)

func task3() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
		wg.Done()
	}()

	wg.Wait()
}

func task4() {
	ch := make(chan int, 5)

	// buffered channels won't block
	//  when sending and not immediately received
	//  therefore, its possible to send in the main thread
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
		wg.Done()
	}()

	wg.Wait()
}
