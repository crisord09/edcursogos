package main

import "fmt"

// || Or
// $$ And
// ! Convertir a su valor contrario
func main() {
	isValid := false

	if isValid {
		fmt.Println("esto esta en el bloque de verdadero")
		//fmt.Printf("%T\n", isValid) // Me avisa el tipo de dato que es
		fmt.Println(isValid)
		if 5 < 10 {
			fmt.Println("5 es mayor que 10")
		}
	} else {
		fmt.Println("Esto esta en el bloque falso")
		fmt.Println(isValid)
	}

	fmt.Println("Fin del Programa")
}
