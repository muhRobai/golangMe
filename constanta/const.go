package main

import (
	"fmt"
)
// explisit
const (
	me = "robai"
	frist = "Muhamad"
)

func main()  {
	command := me
	switch command {
	case me:
		fmt.Println("Call me",me)
	case frist:
		fmt.Println("My Frist Name :", frist)
	default:
		fmt.Println("who are you?")
	}
}