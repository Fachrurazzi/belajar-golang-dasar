package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Banjarmasin", "Kalimantan Selatan", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Banjarbaru"

	*address2 = Address{"Batulicin", "Kalimantan Selatan", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)

	address4.City = "Barabai"
	fmt.Println(address4)
}
