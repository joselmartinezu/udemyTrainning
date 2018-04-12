package main

import "fmt"

func cambiar(valor *int32) {

	*valor = 0

}

func main() {
	var a int32 = 50

	fmt.Println(&a)
	fmt.Printf("%p\n", &a)

	cambiar(&a)

	fmt.Println("El valor de a es: ", a)

}
