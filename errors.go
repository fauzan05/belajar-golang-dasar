package main

import (
	"belajar-golang-dasar/pembagian"
	"fmt"
)

func main(){
	hasil,err := pembagian.Pembagian(0,2)
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(err)
	}
}