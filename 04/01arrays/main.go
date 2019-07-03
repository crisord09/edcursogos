package main

import "fmt"

func main() {
	//Arrays
	//Todos los valores deben de ser el mismo tipo
	//Tamaño fijo
	/*
		var nombres [3]string
		nombres[0] = "Christian"
		nombres[1] = "Edgar"
		nombres[2] = "Blass"
		//nombres[3] = "Pedro"

		fmt.Println(nombres[0])*/
	nombres := [3]string{"Daniel", "Alvaro", "Alexys"}
	//nombre := nombres[1]
	size := len(nombres)

	fmt.Println("El tamaño del arreglo es de:", size)
	fmt.Println(nombres)
	//fmt-println(nombres[inicio:final(excluyente)])
	//fmt.Println(nombres[1:2])
}
