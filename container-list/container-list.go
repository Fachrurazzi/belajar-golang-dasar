package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Razzi")
	data.PushBack("Tirta")
	data.PushBack("Fachrurazzi")

	//for e := data.Front(); e != nil; e = e.Next() {
	//	fmt.Println(e.Value)
	//}

	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

}
