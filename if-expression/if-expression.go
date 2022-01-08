package main

import "fmt"

func main() {
	name := "Kurniawan"

	if name == "Tirta" {
		fmt.Println("Hello " + name)
	} else if name == "Joko" {
		fmt.Println("Hello " + name)
	} else {
		fmt.Println("Hi, Boleh Kenalan ?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
