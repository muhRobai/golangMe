package main

import (
	"fmt"
)

type persegiPajang struct{
	Panjang int
	Lebar int
}

func (Bidang persegiPajang) luas() int{
	return Bidang.Panjang * Bidang.Lebar
}

func (Bidang persegiPajang) Keliling() int{
	return (Bidang.Panjang + Bidang.Lebar) * 2
}

func main(){
	pp := persegiPajang{
		Panjang: 8,
		Lebar: 6,
	}

	fmt.Println("luas bidang: ", pp.luas())
	fmt.Println("Keliling Bidang :", pp.Keliling())
}