package main

import (
	"fmt"
	"time"
	"sync"
)

func Q3(){
	// init new channel
	ch := make(chan int)
	var wg sync.WaitGroup

	// output main function invocation
	fmt.Println("Main Function Executing...")
	time.Sleep(time.Second * 1)

	// create waitGroup to store number in channel
	wg.Add(1)	
	go func(){
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()	
	
	// create WaitGroup to consume number from channel
	wg.Add(1)

	go func(){
		defer wg.Done()
		for num := range ch {
			fmt.Println(num)
			time.Sleep(time.Second * 1)
		}
	}()

	// wait for 
	wg.Wait()

	// have a delay of 2s before outputting exit message
	time.Sleep(time.Second * 2)
	fmt.Println("Main Function Execution Finished...")
}