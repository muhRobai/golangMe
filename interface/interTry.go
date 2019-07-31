package main

import(
	"fmt"
)

type Bidang interface{
	luas() int
	keliling() int
}

func printBidang(b Bidang){
	fmt.Println("luas bidang :", b.luas())
	fmt.Println("keiling bidang :", b.keliling())
	fmt.Println(" ")
}

type PersegiPanjang struct{
	Panjang int
	Lebar int
}

func (Bid PersegiPanjang) luas()int {
	return Bid.Panjang * Bid.Lebar
}

func (Bid PersegiPanjang) keliling() int{
	return (Bid.Panjang + Bid.Lebar) * 2
}

type persegi struct{
	sisi int
}

func (Bid persegi) luas() int{
	return Bid.sisi * Bid.sisi
}

func (Bid persegi) keliling() int{
	return Bid.sisi * 4
}

func main(){
	pp := PersegiPanjang{8,6}
	p := persegi{4}
	printBidang(pp)
	printBidang(p)
}