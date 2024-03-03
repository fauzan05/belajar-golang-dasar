package main

import "fmt"

func getGoodbye(name string) string {
	return "Hello " + name
}

func main(){
	get := getGoodbye

	fmt.Println(get("fauzan"))
}