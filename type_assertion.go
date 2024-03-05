package main

import (
	"fmt"
	"strconv"
)

// type assertion adalah merubah tipe data interface pada function ke tipe data lain
func showMessage() interface{} {
	return "123"
}
func showMessageNew() any {
	return "123"
}
func main() {
	getMessage := showMessage()
	fmt.Printf("Tipe datanya : %T ", getMessage)
	// Mengonversi nilai dari interface{} ke string terlebih dahulu
	strValue, ok := getMessage.(string)
	if !ok {
		fmt.Println("Gagal mengonversi nilai ke string")
		return
	}
	// Mengonversi string ke integer menggunakan strconv.Atoi()
	getMessageResult, err := strconv.Atoi(strValue)
	fmt.Printf("Tipe datanya : %T, valuenya : ", getMessageResult)
	if err != nil {
		fmt.Println("Gagal mengonversi string ke integer:", err)
		return
	}
	fmt.Println(getMessageResult)

	// cara lain
	result := showMessageNew()
	switch value := result.(type) {
	case string:
		fmt.Printf("Tipe datanya : %T", value)
	case int:
		fmt.Printf("Tipe datanya : %T", value)
	default:
		fmt.Println("Tipe datanya : %T", value)
	}



}
