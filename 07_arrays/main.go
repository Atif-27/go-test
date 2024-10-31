package main

import "fmt"

func main() {
	//Numbered sequence of specific length
	var nums [4]int
	// default value are added
	// int -> 0 , bool -> false  , string->  ""
	var names [4] string
	fmt.Println(len(nums))
	nums[0]=1
	nums[1]=2
	nums[2]=3
	nums[3]=4
	fmt.Println(names)

	// ! Decalre and initialise
	age := [3]int{1,2,4}
	fmt.Println(age)

	//! 2d Array
	twoD:= [2][2]int {{3,4},{4,5}}
	fmt.Println(twoD)

}