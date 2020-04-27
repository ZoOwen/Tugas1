package main

import "fmt"

func main(){
	usd := 2
	rp := 15524.50
	floatusd :=float64(usd)
	konversi :=  floatusd * rp
	fmt.Println(usd, "dollar, convert to Rupiah = Rp.",konversi)

}