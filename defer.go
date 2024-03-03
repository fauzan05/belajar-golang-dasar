package main

import "fmt"

func logging() {
	fmt.Println("Fungsi yang dieksekusi terakhir")
	message := recover() // untuk mengambil error yang ada di panic
	fmt.Println("errornya : ", message)
}

func exec() {
	defer logging() // akan diekskusi terakhir
	fmt.Println("Eksekusi pertama")
}

func runApp(error bool){
	defer logging()
	if error {
		panic("Error")
	}
}

func main() {

	runApp(true)
}
