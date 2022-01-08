package main

import "fmt"

func main() {
	const (
		firstName string = "Razzi"
		lastName         = "Tirta"
		value            = 1000
	)

	fmt.Println(firstName + " " + lastName)
	fmt.Println(value)
}
