package main

import (
	"fmt"
	"sync"
)

func task(id int,w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Doing task ",id)
}

func main() {
	var wg sync.WaitGroup
	for i:=0; i<10;i++{
		wg.Add(1)
		go task(i,&wg)
	}
	fmt.Println("Before wait group")
	wg.Wait() // Program waits here , next line not executed untill all waits are removed
	fmt.Println("After Wait group")
}