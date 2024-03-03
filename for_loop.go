package main

import "fmt"

func main(){
	counter := 1

	fmt.Println("For Loop 1 : ")
	for i := counter; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("For Loop 2 : ")
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// foreach dengan index (key)
	names := []string{"fauzan", "susi", "eko", "fajar", "sinta"}
	for index, name := range names {
		fmt.Println("index : ", index, ", value : ", name)
	}
	// foreach tanpa index (key)
	for _, name := range names {
		fmt.Println("value : ", name)
	}
}