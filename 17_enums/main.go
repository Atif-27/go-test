package main

import "fmt"

type orderStatus int
const (
	Confirmed orderStatus = iota
	Delivered
	Cancelled
)
func changeOrderStatus(status orderStatus) {
	fmt.Println("Changed Status to",status)
}


type candidateStatus string
const(
	Accepted candidateStatus = "accepted"
	Rejected ="rejected"
	pending="pending"
)
func changeCandidateStatus (status candidateStatus){
	fmt.Println("Candidate status ",status)
}
func main() {
	changeOrderStatus(Confirmed)
	// changeOrderStatus(Delivered)
	// changeOrderStatus(Cancelled)
	changeCandidateStatus(Accepted)
}