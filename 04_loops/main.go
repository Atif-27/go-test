package main

import "fmt"

func main() {
	//! While loop Implementation
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}
	//! For loop
	for i:=1; i<=5;i++{
		fmt.Println(i)
	}
	//! infinite Loop
	// for {
	// 	fmt.Println("Go is awesome")
	// }
	//! Break and continue in GO
	for i:=1; i<=10;i++{
		if i%2==0{
			continue
		}	
		if i==7 {
			break
		}
		fmt.Println(i)
	}
	//! Range
	for i:=range 10{ // 0 to (range-1)
		fmt.Println(i)
	}
}