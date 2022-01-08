package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`f([a-z])i`)

	fmt.Println(regex.MatchString("fachrurazzi"))
	fmt.Println(regex.MatchString("fachri"))
	fmt.Println(regex.MatchString("fachruRazzi"))

	search := regex.FindAllString("fai fti fauzi fazzi windi", -1)
	fmt.Println(search)
}
