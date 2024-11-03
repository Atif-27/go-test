package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu sync.Mutex
}

func (p *post) increaseCount(w *sync.WaitGroup) {
	defer func(){
		w.Done()
		p.mu.Unlock()
	}()
	p.mu.Lock()
	p.views++
}

func main() {
	var wg sync.WaitGroup
	myPost:= post{
		views: 0,
	}
	for i:=0;i<1000;i++{
		wg.Add(1)
		go myPost.increaseCount(&wg)
	}
	wg.Wait()
	fmt.Println("FInal ans: ", myPost.views)
}