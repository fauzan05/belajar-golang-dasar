package main

import "fmt"

func main(){
	name := "faduzan"

	switch name {
	case "fauzan": 
		fmt.Println("ini fauzan")
	case "susi":
		fmt.Println("ini susi")
	default:
		fmt.Println("tidak ada satupun yang cocok")
	}

	// versi short statement
	switch length := len(name); length < 5{
	case true:
		fmt.Println("ya karakternya kurang dari 5")
	case false:
		fmt.Println("tidak, karakternya lebih dari lima")
	}

	// versi switch tanpa kondisi
	length := len(name)
	switch {
	case length < 5:
		fmt.Println("ya karakternya kurang dari 5")
	case length > 5:
		fmt.Println("tidak, karakternya lebih dari lima")
	}
}