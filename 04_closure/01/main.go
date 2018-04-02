package main

import "fmt"

var x int = 20

func incrementar() int {
	x++
	return x

}

func main() {

	fmt.Println(incrementar())

}
