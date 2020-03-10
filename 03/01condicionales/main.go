package main

import "fmt"

// || Or
// $$ And
// ! Convertir a su valor contrario
// == nos sirve para comparar
func main() {
	/*TABLA DE LA VERDAD
	a	b	a AND(&&) b		a Or b
	v	v		v			  v
	v	f		f			  v
	f	v		f			  v
	f	f		f			  f

	*/
	isValid := false


	
//cuando declaramos una variable en comienzo del bloque if va a estar disponible dentro de ese IF
	if edad, nombre:= 15, "Alvarito"; edad < 14 {
		fmt.Println(nombre, "Es un niño")
	} else if edad < 18 {
		fmt.Println(nombre,"Es un adolecente")	
	} else if edad < 30 {
		fmt.Println("Es un adulto")
	}else{
		fmt.Println("Es un adulto grandecito")
	}
	

	if isValid {
		fmt.Println("esto esta en el bloque de verdadero")
		//fmt.Printf("%T\n", isValid) // Me avisa el tipo de dato que es
		fmt.Println(isValid)
		if 5 < 10 {
			fmt.Println("5 es menor a 10")
		} else{
			fmt.Println("5 no es menor a 10")
		}
	} else {
		fmt.Println("Esto esta en el bloque falso")
		fmt.Println(isValid)
	}

	fmt.Println("Fin del Programa")
}
