package main

import (
	"fmt"
	"time"
)

func main() {
	//! Classic Switch
	var num int = 4
	switch num {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("idk")
	}
	
	//! Multiple Option Switch
	switch time.Now().Weekday() {
		case time.Monday, time.Sunday:
			fmt.Println("Worst day of all time")
		case time.Friday:
			fmt.Println("Friday hai")
		case time.Saturday:
			fmt.Println("Best day of the week")
		default :
			fmt.Println("Invalid Weekday")
	}


	//! Function in switch
	// Functions in Go are first class citizen which means they can be stored in variable and passed as parameter of a function just like in javascript\
	//interface{} means type any
	whoAmI:= func (i interface{})  {
		switch i.(type){
			case int:
				fmt.Println("Its integer")
			case bool:
				fmt.Println("Its Boolean")
			case string:
				fmt.Println("It is a string")
			default:
				fmt.Println("Other")
		}
		
	}
	whoAmI(23.5)
}