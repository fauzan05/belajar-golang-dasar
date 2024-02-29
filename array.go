package main

import "fmt"
var names [3] string
	
func main(){

	// names[0] = "fauzan"
	// names[1] = "susi"
	// names[2] = "rudi"

	// fmt.Println(names[0])
	// fmt.Println(names[1])
	// fmt.Println(names[2])

	var values = [3]int{
		1,
		2,
		3,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 9
	fmt.Println(values)

	// cara agar value dalam array bebas jumlahnya mau berapa indeks
	var unlimited_values = [...]int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
	}
	 
	fmt.Println(unlimited_values)
	fmt.Println(len(unlimited_values))
}