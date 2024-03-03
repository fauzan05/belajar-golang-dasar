package main

import "fmt"

func main(){
	name := "Fauzan"

	if name == "Fauzan" {
		fmt.Println("Ya, namanya Fauzan")
	} else if name == "Susi" {
		fmt.Println("Ya, namanya Susi")
	} else {
		fmt.Println("Tidak ada yang cocok")
	}

	// versi short statement
	if length := len(name); length < 10 {
		fmt.Println("Panjang karakternya kurang dari 10")
	}
}