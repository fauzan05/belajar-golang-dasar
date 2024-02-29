package main

import "fmt"

func main(){
	var nilaiAkhir = 90
	var absensi = 81

	var nilaiLulusAkhir = nilaiAkhir > 80
	var absensiLulus = absensi > 80
	var hasil = nilaiLulusAkhir && absensiLulus
	fmt.Println(hasil) // true
}