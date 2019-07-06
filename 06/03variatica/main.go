package main

import "fmt"

func main() {
	fmt.Println("")
	saludarVarios(20, "Alexys", "Christian", "tery", "Diana")
}

//Pueden ser 1 o varios parametros en eos consiste la funcion variatica
// solo puede tener un unico parametro variatico y debe estar al ffinal de la funcion como nobmres
func saludarVarios(edad uint8, nombres ...string) {
	// Que tipo de de dato es Nombres
	//fmt.Printf("%T\n", nombres)
	for _, v := range nombres {
		fmt.Println("Hola", v, "tienes", edad, "años de edad")
	}
}
