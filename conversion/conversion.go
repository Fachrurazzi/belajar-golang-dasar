package main

import "fmt"

func main() {
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Fachrurazzi"
	var f = name[0]
	var fString string = string(f)

	fmt.Println(name)
	fmt.Println(fString)

}
