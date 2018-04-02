package main

import "fmt"

func main(){

	var x int =200

	valor := func () int {
		x ++
		return x
	}

	fmt.Println(valor())

}
