package main

import "fmt"
import "strings"
import(
	"string/soal1"
)

type student struct{
	name string
	
}

func (s student) sayHello(){
	fmt.Println("hello", s.name)
}

func (s student) getNameAt(i int) string{
	return strings.Split(s.name, " ")[i]
}

func main(){
	// var s1 = student{"jhon wick"}

	// s1.sayHello()
	// var name = s1.getNameAt(0)
	// fmt.Println("nama saya", name)

	// var name = "coba ini"
	// var hasil = strings.Split(name,"")

	// fmt.Printf("ini tipe data %T dari text %v", hasil, hasil)
	soal1.Heloo("omama")
}