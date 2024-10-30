package main

import "fmt"

const global int=100
func main() {
	name := "Docker and Kubernetes are written in Go"
	// name = "Go is slow" ** Cannot change
	fmt.Println(name)

	const (
		PORT=3000
		host="localhost"
	)
	fmt.Println(PORT,host)
}