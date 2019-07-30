package main

import (
	"fmt"
)

func main(){
	numb := 7
	for i := 0; i < numb; i++ {
		for j := numb ; j > i; j-- {
			fmt.Print(" ")
		}
		for j := 0; j <= numb - (numb - i); j++ {
			fmt.Print("-")
		}
		for j := numb ; j > i; j-- {
			fmt.Print("*")
		}
		fmt.Println(" ")
	}

	fmt.Println("===================")

	for i := 0; i < numb; i++ {
		for j := 0; j <= numb - (numb - i); j++ {
			fmt.Print(" ")
		}
		for j := numb ; j > i; j-- {
			fmt.Print("+")
		}
		fmt.Println(" ")
	}

	fmt.Println("======================")

	indexs := map[int]string{
		1: "judul",
		2: "bab 1",
		3: "bab 2",
		4: "bab 3",
	}

	for key, val := range indexs{
		fmt.Println("halaman",key,"adalah", val)
	}


	fmt.Println("======================")
	value := []int{1,2,3,4,5,6,7,8,9}
	
	for i, val := range value{
		fmt.Println("array ke ",i, "adalah ", val)
	}

	


}