package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke-", counter)
	}

	slice := []string{"Eko", "Kurniawa", "Razzi"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, v := range slice {
		fmt.Println("Index", i, "=", v)
	}

	person := make(map[string]string)
	person["name"] = "Razzi"
	person["tittle"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
