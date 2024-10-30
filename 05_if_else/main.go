package main

import "fmt"

func main() {
	//! Classical If Else IF
	age := 20
	if age <= 10 {
		fmt.Println("You are a child")
	} else if age <20{
		fmt.Println("You are a teenager")
	} else{
		fmt.Println("You are an adult")
	}

	//! Logical Cond
	var user="admin"
	var hasPerm= true
	if user=="admin" || hasPerm{
		fmt.Println("You can Access")
	}
	if user=="admin" && hasPerm{
		fmt.Println("You have full access")
	}

	//! Declare variable inside if construct using ;
	if i:=10; i>5{
		fmt.Println("Greater than 5, ",i)
	}else{
		fmt.Println("Less than 5", i)
	}
	//Note- No Ternary operator in GO
}