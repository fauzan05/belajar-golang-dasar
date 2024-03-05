package main

import (
	"fmt"
	"belajar-golang-dasar/database" // ketika package tidak digunakan, gunakan tanda _ sebelum nama package agar tidak eror
)

func main(){
	fmt.Println(database.GetDatabase())
	// fmt.Println("hello")
}