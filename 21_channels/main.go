package main

import (
	"fmt"
	"sync"
	"time"
)



func print(numChan chan int, w *sync.WaitGroup){
	defer  w.Done()
	fmt.Println("Inside Fn Waiting")
	res:= <-numChan
	res2:= <-numChan
	fmt.Println("Result", res)
	fmt.Println("Result2", res2)
}


func main() {
	var wg sync.WaitGroup
	messageChan := make(chan int)
	wg.Add(1)
	go print(messageChan,&wg)
	fmt.Println("Created goroutines now lets sleep 5 sec")
	time.Sleep(time.Second *5)
	messageChan <- 10
	messageChan <- 40
	wg.Wait()
}