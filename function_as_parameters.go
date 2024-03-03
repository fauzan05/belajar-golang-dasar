package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string)string) {
	fmt.Println("Hello ", filter(name))
}

func filterName(name string) string{
	if name == "Anjing" {
		return "Tidak boleh"
	} else {
		return name
	}
}

// membuat alias untuk function parameter
type Filter func(string)string

func sayHelloWithFilterNew(name string, filter Filter) {
	fmt.Println("Hello, ", filter(name))
}

func main(){
	filter := filterName
	sayHelloWithFilter("Anjing", filter)
	sayHelloWithFilterNew("Fauzan", filter)
}