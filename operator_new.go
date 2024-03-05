package main

import "fmt"

type Address struct{
	Kabupaten string
	Provinsi string
	Negara string
}

func main(){
	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.Kabupaten = "Banjarnegara"

	fmt.Println(alamat1) // Banjarnegara
	fmt.Println(alamat2) // Banjarnegara
}