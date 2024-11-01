package main

import (
	"fmt"
	"time"
)
type user struct{
	id string
	name string
}
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	userInfo user
}
//! Constructor Hack
func newOrder(id string,amount float32,status string,createdAt time.Time) *order{
	myOrder:= order{
		id:id,
		amount: amount,
		status: status,
		createdAt: createdAt,
	}
	return &myOrder
}
//! Struct Function
func (o *order) changeOrderStatus(status string){
	o.status=status
}

func main() {
	//! It is not mandatory to provide value for all the fields
	var myOrder order= order{
		id:"2dftopevfejjsncdmfnfexiqjwnfhmwbxen",
		amount: 2000.00,
		status: "Out For Delivery",
		userInfo: user{"900","Atif"},
	}
	myOrder.createdAt= time.Now()
	myOrder.changeOrderStatus("Delivered")
	fmt.Println(myOrder)
	bigBillionOrder:= newOrder("fb6s8kndtzps6wnm53gfk35gfsre",200,"Confirmed",time.Now())
	fmt.Println(bigBillionOrder)

	//! If datatype must be used only once
	userInfo:=struct{
		name string
		address string
		age int
	}{"Atif","123 street",23}
	fmt.Println(userInfo)
}