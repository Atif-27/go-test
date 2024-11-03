package main

import (
	"fmt"
)

func addData(res chan int,a int, b int) {
	fmt.Println("Making result")
	res <- a+b
}

func main() {
	result := make(chan int)
	go addData(result,4,5) 
	fmt.Println("Waiting for result")
	fmt.Println("The result is", <-result) // No need to add wait group as the channel waits for output, which we get after go routine is executed
	fmt.Println("End")
}