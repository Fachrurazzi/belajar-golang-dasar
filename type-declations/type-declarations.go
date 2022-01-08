package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTP NoKTP = "6371031308980005"
	var marriedStatus = true
	fmt.Println(noKTP)
	fmt.Println(marriedStatus)

}
