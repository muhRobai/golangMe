package main

import (
	"fmt"
)

func main (){
	himpunan := [10]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(himpunan)

	himpunan[4] = 24
	fmt.Println("array number 5 : ", himpunan[4])
	fmt.Println(himpunan)

	//add
	name := []string{"muhamad","robai","ganteng"}
	fmt.Println(name)

	// remove
	name = append(name, "tapi")
	fmt.Println(name)

	name = name[:len(name)-1]
	fmt.Println(name)

	indexs := map[string]int{
		"judul":1,
		"bab 1":10,
		"bab 2":25,
		"bab 3":50,
	}
	fmt.Println(indexs)
	fmt.Println(indexs["judul"])

	indexs["bab 5"] = 100
	fmt.Println(indexs)

	delete(indexs, "bab 3")
	fmt.Println(indexs)

}