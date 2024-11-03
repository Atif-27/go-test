package main

import (
	"fmt"
	"time"
)

func task(channelSync chan bool) {
	defer func() { channelSync <- true }()
	time.Sleep(time.Second*5)
	fmt.Println("Processing")
}
func main() {
	taskChannel:= make(chan bool)
	fmt.Println("Start")
	go task(taskChannel)
	fmt.Println("Wait")
	<- taskChannel
	fmt.Println("END")
}