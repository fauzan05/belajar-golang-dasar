package main

import "fmt"

// jika menggunakan slice, berarti []int, bukan ...int
func sumAll(name string,numbers ...int) (string, int){
	total := 0
	for _, number := range numbers {
		total += number
	}
	return name + " totalnya adalah : ", total
}

func main(){
	fmt.Println(sumAll("Fauzan", 10,10,10,10,10))
	// fmt.Println(sumAll("Fauzan", []int{ 10,10,10,10,10})) // jika menggunakan slice

	numbers := []int{1,2,3,4,5,6,7,8,9,10,11}
	fmt.Println(sumAll("fauzan", numbers...)) // jika ingin menggunakan array yang sudah ada, maka tambahkan...
}