//el recover se usa dentro de las funciones defer
package main

import "fmt"

func main() {
	f()

}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%T\n", r)
			fmt.Println("Recuperado en f")
		}
	}()

	fmt.Println("LLamando g.")
	g(5)
	fmt.Println("Retorna normalmente desde g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("AAAAAAAAAAAaaaaa")
		panic("el número no puede ser mayor que 3")
	}
	defer fmt.Println("Defer en la función g")
	fmt.Println("Imprimiendo en g", i)
}
