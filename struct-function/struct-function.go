package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHi() {
	fmt.Println("Hallo", a.Name)
}

func main() {
	var cust Customer
	cust.Name = "Razzi"
	cust.Address = "Banjarmasin"
	cust.Age = 24

	cust.sayHello("Razzi")
	cust.sayHi()

	// fmt.Println(cust)

	// var razzi Customer = Customer{
	// 	Name:    "Razzi",
	// 	Address: "Banjarmasin",
	// 	Age:     24,
	// }
	// fmt.Println(razzi)

}
