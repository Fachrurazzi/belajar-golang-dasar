package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var cust Customer
	cust.Name = "Razzi"
	cust.Address = "Banjarmasin"
	cust.Age = 24

	fmt.Println(cust)

	var razzi Customer = Customer{
		Name:    "Razzi",
		Address: "Banjarmasin",
		Age:     24,
	}
	fmt.Println(razzi)

}
