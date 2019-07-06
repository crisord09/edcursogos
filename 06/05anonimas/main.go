package main

import "fmt"

func main() {
	//las funciones anonimas son un tipo de dato pude asignarse un tipo de dato

	//Función con parentesis al final se autoejecuta
	/*func() {
		fmt.Println("Me imprimo")
	}()*/

	//Nos permite usar una función como una variable
	/*anonima := func() {
		fmt.Println("Me imprimo estando en una variable llamada anonima")
	}

	anonima()*/

	secuencia1 := secuencia()
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())

	fmt.Println("")

	secuencia2 := secuencia()
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
}

//Closhure o función anonima retorna otra función
func secuencia() func() int32 {
	var x int32
	return func() int32 {
		x++
		return x
	}
}
