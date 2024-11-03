package main

import (
	"fmt"
	"time"
)

func emailSender(emailChan <-chan string, doneChan chan<- bool) {
	defer func ()  {
		doneChan<-true
	}()
	for i:= range emailChan{
		time.Sleep(time.Second)
		fmt.Println("Sending Email to: ",i)
	}
}
func main() {
	emailChan := make(chan string, 100)
	doneChan:= make(chan bool)
	go emailSender(emailChan,doneChan)
	for i:=0; i<10;i++{
		emailChan<- fmt.Sprintf("%d@gmail.com",i)
	}
	fmt.Println("All Scheduled")
	close(emailChan)
	<- doneChan
}