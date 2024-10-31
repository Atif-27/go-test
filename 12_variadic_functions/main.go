package main

import "fmt"

func main() {
	res := sum(2, 4, 5, 6, 67, 7, 8, 91, 23, 1)
	fmt.Println(res)
	//in case of slice
	sl:= []int {1,23,45,4,23234,3}
	sum(sl...)
}

func sum(val ...int) int {
	total := 0
	for _, val := range val {
		total = total + val
	}
	return total
}