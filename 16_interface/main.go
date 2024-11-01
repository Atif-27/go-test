package main

import "fmt"

type paymenter interface{
	pay(amount float32)
}

type payment struct{
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	//! Everytime we need to add new Payment gateway we need to update fn which is against SOLID principles -> Open for extension close for modification
	p.gateway.pay(amount)
}

//! Razorpay
type razorpay struct{
	id string
}
func (r razorpay) pay(amount float32) {
	fmt.Println("Making Payment using razorpay",amount)
	//Api calls
}
//! Stripe
type stripe struct{
	id string
}
func (s stripe) pay(amount float32){
	fmt.Println("Making payment using stripe",amount)
}

func main() {
	stripeGw:= stripe{
		id:"@3232323",
	}
	razorpayGw:= razorpay{
		id:"@3232323",
	}
	
	myPayment:= payment{
		gateway: razorpayGw,
	}
	myPayment.makePayment(100)
	
	myPaymen2t:= payment{
		gateway: stripeGw,
	}
	myPaymen2t.makePayment(20)
}

// Everytime u need to change gateway it is needed to be changed in struct
// You cannot test the payment coz we cannot pass a dummy payment gateway (fake implementation)

// The thing with interface is only the struct which implement the interface has the type of interface -> so we can keep it generic for all instance of different struct