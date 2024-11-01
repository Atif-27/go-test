package main

import (
	"fmt"
	"math"
)

type shaper interface{
	getArea() float32
}

type circle struct {
	radius float32
}
func (c circle) getArea() float32 {
	return float32(math.Pi * math.Pow(float64(c.radius),2))
}
type rectange struct{
	length float32
	breadth float32
}
func (r rectange) getArea() float32{
	return r.length*r.breadth
}


func printArea(s shaper){
	fmt.Println(s.getArea())
}

func main(){
	c1:=circle{
		radius: 2,
	}
	r1:=rectange{
		length: 10,
		breadth: 20,
	}
	printArea(c1)
	printArea(r1)
}