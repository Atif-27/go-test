package main

import (
	"fmt"
	"maps"
)

// maps are associative tables
func main(){
	//! Maps
	m:= make(map[string]string)
	m["name"]="Atif"
	m["address"]="123 street"
	fmt.Println(m["address"])
	//! If a key is not present then it return zero value (default)
	fmt.Println(m["phone"])
	fmt.Println(m)
	fmt.Println("Length",len(m))
	delete(m,"address")
	fmt.Println(m)
	clear(m)
	fmt.Println(m)


	// ! Conventional way to access map item
	store:= map[string]int{"amount":100,"Stocks":2}
	home:= map[string]int{"amount":100,"Stocks":3}
	val, ok := store["amount"]
	if(!ok){
		fmt.Println("Not Okay")
	} else{
		fmt.Println("All Good the value is ",val)
	}
	fmt.Println(maps.Equal(store,home))

}