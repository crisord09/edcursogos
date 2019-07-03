package main

import "fmt"

func main() {

	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println(" ")
	/*
		for j := 4; j >= 1; j-- {

			//nos permite saltar una iteración
			if j == 3 {
				continue
			}
			fmt.Println(j)
		}
	*/
	/*
		for j := 0; j < 5; j++ {
			if j == 2 {
				fmt.Println("j vale 2 y se rompe el ciclo")
				break
			}
			fmt.Println(j)
		}
	*/
	// var matriz [3][3]int
	matriz := [3][3]int{}
	valor := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			valor++
			matriz[i][j] = valor
		}
	}

	for fila := 0; fila < 3; fila++ {
		fmt.Println(fila)
		for columna := 0; columna < 3; columna++ {
			fmt.Println(matriz[fila][columna])
		}
	}
}
