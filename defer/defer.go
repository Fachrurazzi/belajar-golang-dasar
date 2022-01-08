package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplicaion(value int) {
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApplicaion(0)
}
