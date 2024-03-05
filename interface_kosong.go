package main

import "fmt"

// membuat method dengan return interface kosong
// tipe data interface = any
func Ups() interface{} {
	return true
}

func benar() bool{
	return true
}

func main(){
	kosong := Ups()
	fmt.Println(kosong)
	fmt.Println(benar())
} 