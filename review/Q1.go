package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func printLetters(){
	letters := "abcdefghij"
	for _, letter := range letters {
		fmt.Printf("%c\n",letter)
		time.Sleep(1 * time.Second)	
	} 
}

func Q1(){
	fmt.Println("Begin main function execution...");
	time.Sleep(2 * time.Second) 

	go printNumbers()	
	go printLetters()
	
	time.Sleep(2 * time.Second)
	fmt.Println("Main program execution ends")
}