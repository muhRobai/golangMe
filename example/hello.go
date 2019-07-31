package main

import (
	"fmt"
)

func main(){

	// println or print for string
	fmt.Println("hello")
	text := "this is a string"
	fmt.Println(text)
	var textAlso string = "this is also a text"
	fmt.Print(textAlso)

	// printf for string combine integer
	a := 9999
	fmt.Printf("an Integer with default type %T: %v\n", a, a)

}