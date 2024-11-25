package main

import (
	"fmt"
	"sync"
	"time"
)

func Q5(){
	// declare even and odd channels
	even := make(chan int)
	odd := make(chan int)

	// create waitGroup to wait for both filling
	// process to complete
	var wg sync.WaitGroup

	fmt.Println("Begin main program execution for Q5...")

	// create a waitGroup for evenNumbers filling
	wg.Add(1)
	go func(){
		defer wg.Done()	
		
		for i := 1; i <= 20; i++ {
			// fill odd channel with is
			odd <- i	
		}
		
		close(odd)
	}()

	// create a waitGroup for oddNumbers filling
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		
		for i := 2; i <= 20; i++ {
			// fill even channel with is
			even <- i
		}

		close(even)
	}()
	
	// Consume values from both channels until they are closed
	for {
		select {
		case odd_num, ok := <-odd:
			if ok {
				fmt.Println("Received an odd number:", odd_num)
			}
		case even_num, ok := <-even:
			if ok {
				fmt.Println("Received an even number:", even_num)
			}
		default:
			// Break the loop if both channels are closed
			if len(odd) == 0 && len(even) == 0 {
				return
			}
		}

		// Simulate work
		time.Sleep(time.Second * 1)
	}
}