package main

import "fmt"
var nilai32 int32 = 55555
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)	
func main(){
	

	var name = "Fauzan"
	var e = name[0]
	var getString = string(e) //jika tidak dikonversi menjadi string, maka akan menampilkan karakter ASCII

	fmt.Println(getString)

	var desimal float64 = 4.55
	// var bulat int = 5
	var konversi_ke_bulat int = int(desimal)
	var konversi_ke_desimal float64 = float64(desimal * desimal)

	fmt.Println(konversi_ke_bulat)
	fmt.Println(konversi_ke_desimal)
	
	// fmt.Println(nilai32)
	// fmt.Println(nilai64)
	// fmt.Println(nilai16)
}