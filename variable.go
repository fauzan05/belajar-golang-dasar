package main

import "fmt"
// tidak akan error karena tidak dieksekusi di dalam main/dipanggil didalam main
var (
	firstname = "Fauzan"
	lastname  = "Nurhidayat"
)

var x, y string = "Hello", "World" 


func main() {
fmt.Printf("hai")
fmt.Println(x,y)
// akan error karena tidak digunakan
// var (
// 	firstname = "Fauzan"
// 	lastname  = "Nurhidayat"
// )
}
