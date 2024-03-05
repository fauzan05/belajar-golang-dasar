package main

import(
	"belajar-golang-dasar/helper"
	"fmt"
)

func main(){
	getHelper := helper.SayHello("fauzan")
	fmt.Println(getHelper)
	fmt.Println(helper.Application)
} 