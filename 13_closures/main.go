package main

import "fmt"

func main() {
	incrementFn := test()
	fmt.Println(incrementFn())
	fmt.Println(incrementFn())
	fmt.Println(incrementFn())
}

func test() func() int {
	val := 0
	increment := func() int {
		val++
		return val
	}
	return increment
}