package main

import (
	"fmt"
	"maps"
)

func main(){
	person := map[string]string{
		"name" : "fauzan",
		"address" : "kebumen",
	}

	fmt.Println(person)
	fmt.Println(person["name"]) // fauzan
	fmt.Println(person["address"]) // kebumen
	fmt.Println("Length : ",len(person))
	delete(person, "name")
	fmt.Println(person)

	car := make(map[string]string)
	car["model"] = "tesla"
	car["year"] = "2012"
	fmt.Println(car)

	ferrari := map[string]string{
		"model" : "ferrari",
		"year" : "2023",
	}

	tesla := map[string]string{
		"model" : "tesla",
		"year" : "2023",
	}

	if !maps.Equal(ferrari, tesla) {
		fmt.Println("Isi array tidak sama")
	}

}