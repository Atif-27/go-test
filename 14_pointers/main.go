package main

import "fmt"

func main() {
	num:=10
	fmt.Println("Before change ", num)
	changeNum(&num)
	fmt.Println("After change ", num)

}

func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum Func ", *num)
}