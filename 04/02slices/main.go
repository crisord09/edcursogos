package main

import "fmt"

func main() {
	//Slices representan un tamaño de variales del mismo tipo tienen un apuntador al Slice
	//1. Apuntador a un array
	//2. Tamaño
	//3. Capacidad

	/*var nombres []string
	nombres = append(nombres, "Daniel")
	fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Alvaro")
	fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Alexys")
	fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Pedro")
	fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Juan")
	fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))

	nombres[3] = "Juancho"*/

	//make(tipo de dato del slice, tamaño inicial, capacidad inicial(opcional))
	//nombres := make([]string, 0)

	nombres := []string{
		"alvaro",
		"Daniel",
		"Alexys",
	}

	/*
		fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Alvaro")
		fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Alexys")
		fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Pedro")
		fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Juan")
		fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Yeymi")
		fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))

		nombres[3] = "Juancho"
	*/

	fmt.Println(nombres)
}
