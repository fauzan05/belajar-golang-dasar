package main

import ("fmt")

type Address struct{
	Kabupaten string
	Provinsi string
	Negara string
}



func main(){
	address1 := Address{"Kebumen", "Jawa Tengah", "Indonesia"}
	address2 := &address1 // pass by reference
	address2.Kabupaten = "Cianjur"
	fmt.Println(address1)


	x := 10
	ptr := &x
	*ptr = 20
	fmt.Println(x)
	fmt.Println(&ptr)
	fmt.Println(&x)
}