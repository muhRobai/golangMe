package main

import (
	"fmt"
)

func main(){

	// bolean
	 isHuman := true
	 isHandsome := false

	fmt.Println("isHuman ?", isHuman)
	fmt.Println("the human is Hansome ?", isHandsome)

	// numberik

	a := 9999
	fmt.Printf("an Integer with default %T: %v\n", a, a)

	// var with int8 have a max value 127 - (-128)
	var b1 int8 = 127 
	var b2 int8 = -128
	fmt.Println("an 8 bit integer: ", b1)
	fmt.Println("an 8 bit Integer: ", b2)

	// var with uint 8 is only positive value 0 - 255
	var b3 uint8 = 0
	fmt.Println("an 8 bit unsigned integer: ", b3)

	//Float
	d1 := 256.0
	fmt.Printf("A float with default type %T: %v\n", d1,d1)
	var d2 float32 = 0.000001
	fmt.Printf("A float with 32 default type %T: %v\n", d2,d2)

	var d3 float64 = 1000
	fmt.Printf("A float with 64 default type %T: %v\n", d3,d3)

	// Complex number
	sum := 1.6 + 1.2i
	fmt.Printf("A Complex number with default type %T: %v\n", sum, sum)

	var e1 complex64 = 4 + 3i
	fmt.Printf("A Complex number with default type %T: %v\n", e1, e1)

	var e2 complex128 = 4+3i
	fmt.Printf("A complex number with default type %T: %v\n", e2,e2)

	//Byte 
	var f1 byte = 255
	fmt.Printf("A byte with default type %T: %v\n", f1,f1)
	f1--
	fmt.Printf("A byte with default type %T: %v\n", f1,f1)

	//Rune
	var g1 rune = '�'
	fmt.Printf("A rune with default type %T: %U\n", g1, g1)
	fmt.Println(rune(g1))
	

}