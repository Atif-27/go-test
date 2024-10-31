package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func getName() (string, string, bool) {
	return "Test", "Rest", true
}

func process(fn func(a, b int) int) int {
	val := fn(20, 4)
	return val + 10*200
}

func main() {
	res := add(4, 5)
	fmt.Println(res)
	str1,_,str3 := getName()
	fmt.Println(str1,str3)
	result := process(add)
	fmt.Println(result)
}