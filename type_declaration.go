package main

import "fmt"

func main(){
	// membuat sebuah alias tipe data, contohnya membuat alias tipe data string
	type iniString string
	type iniFloat64 float64
	var iniAdalahString iniString = "haha" //jika diisi integer maka akan error, dan penamaan variabelnya tidak boleh sama dengan tipe datanya
	var iniDesimal iniFloat64 = 4.5 //jika diisi integer maka akan error

	fmt.Println(iniAdalahString)
	fmt.Println(iniDesimal)
}