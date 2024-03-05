package main

import (
	"fmt"
)

type Address struct{
	Kabupaten string
	Provinsi string
	Negara string
}

func main(){
	address1 := Address{"Kebumen", "Jawa Tengah", "Indonesia"}
	address2 := &address1
	address2.Kabupaten = "Cianjur"
	fmt.Println(address1)
	fmt.Println(address2)

	// cara seperti ini berarti address2 sudah tidak mengacu ke address1 jika address2 mengalami perubahan
	// address2 = &Address{"Cikarang", "Jawa Barat", "Indonesia"}
	// fmt.Println(address1)
	// fmt.Println(address2)

	// cara seperti ini berarti address1 ikut berubah mengikuti perubahan yang dilakukan oleh address2 menggunakan operator asterisk (*)
	*address2 = Address{"Cakung", "Jakarta Timur", "Indonesia"}
	fmt.Println(address1) // cakung jakarta timur indonesia
	fmt.Println(address2) // cakung jakarta timur indonesia

}