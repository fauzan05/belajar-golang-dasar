package main

import "fmt"

func main(){
	newSlice := make([]string, 2, 5) 
	newSlice[0] = "fauzan"
	newSlice[1] = "susi"
	
	newSlice = append(newSlice, "rudi")
	newSlice = append(newSlice, "heri")
	newSlice = append(newSlice, "iqbal")
	newSlice = append(newSlice, "lisa")
	
	fmt.Println(newSlice) // [fauzan, susi, rudi, heri, iqbal, lisa]
	fmt.Println(len(newSlice)) // [total indeksnya 6]
	fmt.Println(cap(newSlice)) // jika indeks array totalnya lebih dari 5, maka cap akan bernilai 10 karena sudah overkapasitas. Jika dibawah 5 atau sama dengan 5, maka cap akan masih bernilai 5

	// copy slice
	fromSlice := newSlice[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice)) // kosong [  ]

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice) // isinya sama dengan fromSlice karena sudah di copy
}