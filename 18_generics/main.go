package main

import "fmt"

func displaySlice[T comparable, F string|int](arr []T, name F) {
	var sliceName F=name
	for _, val := range arr {
		fmt.Println(val,sliceName)
	}
}

type stack[T int|string] struct{
	elements []T
}
func (s *stack[T]) push(val T){
	s.elements = append(s.elements, val)
}
func (s *stack[T]) pop()T{
	size:=len(s.elements)
	if size == 0 {
		var zero T
		return zero
	}
	val:= s.elements[size-1]
	s.elements=s.elements[:size-1]
	return val
}

func main() {
	iarr:=[]string{"1","2","45"}
	displaySlice[string,string](iarr,"Atif")
	myStack:= stack[int]{}
	myStack.push(10)
	myStack.push(20)
	myStack.pop()
	myStack.push(30)
	myStack.push(40)
	myStack.pop()
	fmt.Println(myStack.elements)
}