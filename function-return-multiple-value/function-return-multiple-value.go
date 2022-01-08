package main

import "fmt"

func getFullName() (string, string, int) {
	return "Razzi", "Tirta", 24
}

func main() {
	firstName, lastName, _ := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}
