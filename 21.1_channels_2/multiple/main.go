package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		chan1 <- 12
	}()
	go func() {
		chan2 <- "Skibidi"
	}()
	for i:=0 ; i<2;i++{
		select{
		case chan1Val:= <- chan1:
			fmt.Println("Received Data from Chan 1", chan1Val)
		case chan2Val:= <- chan2:
			fmt.Println("Received Data from Chan 2", chan2Val)
		}
	}
}