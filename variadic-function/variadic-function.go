package main

import "fmt"

func sumAll(name string, numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total = total + number
	}

	return total
}

func main() {
	total := sumAll("Razzi", 90, 90, 60, 70)
	fmt.Println(total)

	slice := []int{10, 30, 40, 80}
	total = sumAll("Tirta", slice...)
	fmt.Println(total)
}
