//Un puntero es la direccion en memoria a una variable de un tipo determinado valor en memoria es como una variable estatica

package main

import "fmt"

func main() {
	//El valor 0 de un puntero es 1000
	/*a := 0
	fmt.Println(&a)
	*/
	a := 4
	fmt.Println("Antes de duplicar, a =", a)
	//el &me ayuda a ingresar al valor en memoria (Puntero)
	duplicar(&a)
	fmt.Println("despues de duplicar, a =", a)

}

//Debi agregar * para poder explicar que GO debe de usar el valor de un puntero
func duplicar(a *int) {
	*a = *a * 2
	fmt.Println("Dentro de la función duplicar a =", *a)
}
