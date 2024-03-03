package main

import "fmt"

func main(){

	fmt.Println("break : ")
	// break
	for i := 1; i < 10; i++ {
		if(i == 5) {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("continue : ")
	//continue
	for i := 0; i < 10; i++ {
		if(i%2 == 0){
			continue // akan di skip
		}
		fmt.Println(i)
	}
}