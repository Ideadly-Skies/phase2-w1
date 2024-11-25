package main

import (
	"fmt"
	"time"
)

func Q5(){
	// declare even and odd channels
	even := make(chan int)
	odd := make(chan int)

	// create waitGroup to wait for both filling
	// process to complete

	fmt.Println("Begin main program execution for Q5...")

	go func(){
		for i := 1; i <= 20; i += 2 {
			// fill odd channel with is
			odd <- i	
		}
		
		close(odd)
	}()

	go func ()  {
		for i := 2; i <= 20; i += 2 {
			// fill even channel with is
			even <- i
		}

		close(even)
	}()
	
	for {
      select {
		case evenNum, ok := <- even:
			if ok {
				fmt.Println("Even Number:", evenNum)
			} else {
				even = nil
			}
		case oddNum, ok := <- odd:
			if ok {
				fmt.Println("Odd Number:", oddNum)
			} else {
				odd = nil
			}
			time.Sleep(1 * time.Second)
		}

		if even == nil && odd == nil {
			break
		}
   }

	time.Sleep(time.Second * 1);
	fmt.Println("finished main Q5 invokation...");
}