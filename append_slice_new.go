package main

import "fmt"

func main(){
	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu_New"
	daySlice1[1] = "Minggu_New"
	// newSlice := make([]string, 2, 5)
	// copy(newSlice, daySlice1)
	fmt.Println(days) // [senin, selasa, rabu, kamis, jumat, Sabtu_New, Minggu_New]
	fmt.Println(daySlice1)
	// fmt.Println(newSlice)
	
	daySlice2 := append(daySlice1, "libur_baru")
	daySlice2[0] = "ups_masuk" 
	daySlice2[1] = "ups_masuk_lagi" 
	fmt.Println(daySlice2) // [ups_masuk, Sabtu_New, Minggu_New] karena array days sudah habis indeksnya, maka akan membuat array baru didalam variabel daySlice2
	fmt.Println(days) // [senin, selasa, rabu, kamis, jumat, Sabtu_New, Minggu_New]

	daySlice3 := append(daySlice2, "welcome") // menambahkan value baru ke array daySlice2
	fmt.Println(daySlice3)
	fmt.Println(daySlice2)
	fmt.Println(days)
}