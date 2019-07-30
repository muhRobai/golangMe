package main

import (
	"fmt"
)

func main(){
	types := "helow"

	switch types {
	case "get in":
		fmt.Println("you Get in")
	case "helow":
		fmt.Println("hello you")
	default:
		fmt.Println("oke you need me")
	}
}