//comenatrios de linea para documentar en codigo
/*COMENTARIOS
DE
BLOQUE
*/
//Una constante no puede cambiar de dato en toda la app
//

package main

import "fmt"

func main() {
	//const nombre = "Pedro"

	//fmt.Println(nombre)

	var a int
	var b int8
	n := "Pedro"
	p := "Gómez"

	a = 121212
	b = 5

	/*SE REALIZA POR QUE 
	casting int() poder realizar el cambio temporal de una variable para realizar una funcion especifica*/
	c := a + int(b)

	/* Printf le decimos que estamos enviando un dato string permite mostrar el formato que quere
	*/
	fmt.Println(c)
	fmt.Printf("hola %s %s %d ¿cómo estas? \n", n, p, b)
	/*%T nos ayuda a mostrar el tipo de valor de una variable*/
	fmt.Printf("C es de tipo: %T\n", c)

	//prioridad aritmética
	// * / + -
	//modulo es el residuo de la division entre 2 elementos
	// () % * / + -
	d := 6 + 6*(6-6)
	fmt.Println(d)

	// Anque yo no les haya dado un valor a una variable todas comienzan en cero
	//var nombre string
	//var numero int
	//var entiendes bool
	//fmt.Println(nombre, numero, entiendes)

}
