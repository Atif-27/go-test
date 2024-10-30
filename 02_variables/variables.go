package main

import "fmt"

func main() {
	//! String
	var name string = "Go is open-source lang developed by google"
	var address ="123 street" // Typed Infered from the value
	fmt.Println("Name=",name+" Address=", address)
	//! Integer
	var age int=100
	var year=2024
	fmt.Println("Age=",age," Year=",year)
	//! Shorthand
	randomNumber:=20
	var randomString string
	randomString="GoLang"
	fmt.Println("Random Number",randomNumber," Random String:",randomString)
}