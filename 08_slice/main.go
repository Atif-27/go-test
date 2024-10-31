package main

import (
	"fmt"
	"slices"
)

func main() {
	//! CLassic Slice
	var ages []int
	fmt.Println(ages,len(ages), ages==nil)
	// ! Better way to declare slice
	//capacity is the maximum element that can be put
	var names =make([]int,0,5)
	names = append(names, 1)
	names = append(names, 2)
	names = append(names, 3)
	names = append(names, 4)
	//! Copy
	var test= make([]int,len(names),5)
	copy(test,names)
	fmt.Println(test, names)

	
	// ! Slice operator  m to n-1
	fmt.Println(names[1:3]) 
	fmt.Println(names[:3])
	fmt.Println(names[1:])

	//! Slices Package
	fmt.Println(slices.Equal(names,test))
	
	//! 2d slices
	twoD:= [][]int {{2,3},{23,2}}
	fmt.Println(twoD)
}