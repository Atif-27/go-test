package main

import "fmt"

func main() {
	//! Iterating Slices
	nums := []int{2, 3, 45}
	for idx, val := range nums {
		fmt.Println("position",idx,": ",val)
	}

	//! Iterating Maps
	m:= map[string]int {"age":10,"year":2024,"degrees":4}
	for key,val:=range m{
		fmt.Println(key," : ",val)
	}

	//! Iterating Strings
	str:= "Hello World"
	for i, c := range str{
		fmt.Println(i," : ",string(c))
	}
}