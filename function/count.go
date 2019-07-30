package main

import (
	"fmt"
)

func main(){
	sisi := 10
	panjang := 5
	luas := luasPersegi(sisi, panjang)
	keliling := kelilingPersegi(sisi)

	fmt.Println("luas dari Persegi dengan sisi", sisi, "adalah", luas)
	fmt.Println("keliling dari Persegi dengan sisi", sisi, "adalah", keliling)
	


}

func luasPersegi(s int, p int)int{
	luas := s * p
	return luas
}

func kelilingPersegi (s int) int{
	keliling := s * 4
	return keliling
}