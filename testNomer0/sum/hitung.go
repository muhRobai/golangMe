package sum

import (
	"fmt"
	"strings"
	
)

func Sum(text string){
	var role =[]string{"a","i","u","e","o","A","I","U","E","O"}
	var exceptRole = "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"
	var rest = 0
	var exceptRest = 0
	var sumExceptRest = 0
	var txtArray = strings.Split(exceptRole,"")
	// fmt.Println(txtArray[1])

	for j := 0; j < len(role); j++ {
		if strings.ContainsAny(text, role[j]) {
			rest++
		}
	}

	for i := 0; i < len(txtArray); i++ {
		if strings.ContainsAny(text, txtArray[i]) {
			exceptRest = strings.Count(text,txtArray[i])
			fmt.Printf("jumlah dari string %v adalah %v\n", txtArray[i],exceptRest)
			sumExceptRest += exceptRest 
		}
	}

	fmt.Printf("kata %v mempunyai huruf hidup %v dan huruf mati %v\n", text, rest, sumExceptRest)



}