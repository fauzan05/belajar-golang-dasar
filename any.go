package main

import "fmt"

func Ups1() any {
	return 1
}

func count1(a,b int) any{
	return a * b
}

func main(){
	var kosong any = Ups1()
	fmt.Println(kosong)
	fmt.Println(count1(1,2))
}