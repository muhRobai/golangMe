
package main

import (
	"fmt"
)

const (
	me = iota
	frist
)

func main()  {
	command := frist
	switch command {
	case me:
		fmt.Println("Call me",me)
	case frist:
		fmt.Println("My Frist Name :", frist)
	default:
		fmt.Println("who are you?")
	}
}