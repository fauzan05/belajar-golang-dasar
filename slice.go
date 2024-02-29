package main

import "fmt"

func main(){
	names := [...]string{"fauzan", "susi", "budi", "rudi", "heri"}
	
	slice1 := names[1:2] // ambil data dari indeks 1 hingga indeks sebelum 2 (indeks 2 tidak termasuk)
	fmt.Println(slice1)
	fmt.Println("Panjang slice1 = ", len(slice1))

	slice2 := names[:2]
	fmt.Println(slice2)
	fmt.Println("Kapasitas slice2 = ", cap(slice2))

	slice3 := names[2:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)


	// cara lain membuat sebuah slice
	var slice5 []string = names[:]
	fmt.Println(slice5) // hasilnya sama saja dengan slice 4

}