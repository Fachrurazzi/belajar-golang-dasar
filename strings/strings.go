package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Fachrurazzi", "azzi"))
	fmt.Println(strings.Split("Fachru razzi", " "))
	fmt.Println(strings.ToLower("Fachrurazzi"))
	fmt.Println(strings.ToUpper("Fachrurazzi"))
	fmt.Println(strings.ToTitle("Fachrurazzi"))
	fmt.Println(strings.Trim("i    Fachrurazzi       i", " "))
	fmt.Println(strings.ReplaceAll("Razzi Razzi Razzi Andi", "Razzi", "Fachrurazzi"))
}
