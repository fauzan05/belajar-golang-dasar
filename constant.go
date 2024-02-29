package main

import "fmt"

// const adalah variabel yang tidak dapat diubah-ubah
const firstname string = "fauzan"
const lastname = "nurhidayat"

var class = "system information"

func main() {
	// firstname = "haha"
	fmt.Printf("%v", class)

	const Truth = true
	// fmt.Printf("Go rules?", Truth)

}
