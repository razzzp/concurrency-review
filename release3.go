package main

import (
	"fmt"
)

func task5() {
	sendEven := make(chan int)
	sendOdd := make(chan int)

	// send numbers
	go func() {
		for i := 1; i <= 20; i++ {
			if i%2 == 0 {
				//even
				sendEven <- i
			} else {
				sendOdd <- i
			}
		}
		close(sendEven)
		close(sendOdd)
	}()

	// receive
	for {
		select {
		case e, ok := <-sendEven:
			if !ok {
				sendEven = nil
			} else {
				fmt.Printf("Received an even number: %d\n", e)
			}
		case o, ok := <-sendOdd:
			if !ok {
				sendOdd = nil
			} else {
				fmt.Printf("Received an odd number: %d\n", o)
			}
		}
		if sendEven == nil && sendOdd == nil {
			break
		}
	}
}

func taskBonus() {
	sendEven := make(chan int)
	sendOdd := make(chan int)
	sendErr := make(chan error)

	// send numbers
	go func() {
		for i := 1; i <= 22; i++ {
			if i > 20 {
				sendErr <- fmt.Errorf("number %d is greater than 20", i)
				continue
			}
			if i%2 == 0 {
				//even
				sendEven <- i
			} else {
				sendOdd <- i
			}
		}
		close(sendEven)
		close(sendOdd)
		close(sendErr)
	}()

	// receive
	for {
		select {
		case e, ok := <-sendEven:
			if !ok {
				sendEven = nil
			} else {
				fmt.Printf("Received an even number: %d\n", e)
			}
		case o, ok := <-sendOdd:
			if !ok {
				sendOdd = nil
			} else {
				fmt.Printf("Received an odd number: %d\n", o)
			}
		case e, ok := <-sendErr:
			if !ok {
				sendErr = nil
			} else {
				fmt.Printf("Error: %v\n", e)
			}
		}
		if sendEven == nil && sendOdd == nil && sendErr == nil {
			break
		}
	}
}
