package main

import (
	"fmt"
	"sync"
	"time"
)

func Q2(){
	// begin main function execution
	fmt.Println("Begin main function execution...");

	// create a waitGroup for the Goroutine
	var wg sync.WaitGroup 

	// add one goroutine to the waitGroup
	wg.Add(1)

	// print numbers first
	go func(){
		defer wg.Done()
		printNumbers()
	}()

	wg.Wait() // wait for number printing to conclude
	
	// add another goroutine to the waitGroup
	wg.Add(1)

	// print letters next
	go func(){
		defer wg.Done()
		printLetters()
	}()
	
	wg.Wait() // wait for letter printing to conclude

	time.Sleep(2 * time.Second)
	fmt.Println("Main Function Invocation Finished...")
}