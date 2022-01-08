package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2020, 1, 6, 7, 42, 40, 40, time.UTC)
	fmt.Println(utc)

	layout := "2022-01-06"

	format, _ := time.Parse(layout, "2022-01-06")
	fmt.Println(format)
}
