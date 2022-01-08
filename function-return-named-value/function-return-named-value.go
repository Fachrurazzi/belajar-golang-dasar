package main

import "fmt"

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "Razzi"
	middleName = "Tirta"
	lastName = "RTY"

	return
}

func main() {
	fName, mName, lName := getFullName()
	fmt.Println(fName)
	fmt.Println(mName)
	fmt.Print(lName)
}
