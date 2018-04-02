package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {

	var meter float64
	fmt.Println("Por favor introduza la cantidad de metros: ")
	fmt.Scan(&meter)
	yards := meter * metersToYards
	fmt.Println("El cálculo es el siguiente: ", yards)

}
