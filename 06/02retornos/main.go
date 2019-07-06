package main

import "fmt"

func main() {
	/*
		var n1, n2 int8
		n1 = 15
		n2 = 3
		r := suma(n1, n2)

		fmt.Println("La suma de los 2 numeros es:", r)
	*/
	/*var edad uint8
	edad = 17
	fmt.Println(tipoEdad(edad))*/

	n := []int8{52, 63, 48, 5, 5, 3, 7, 100, 2, 47, -5}
	maximo, minimo := maxymin(n)
	fmt.Println("Este es el numero maximo", maximo, "Este es el valor minimo", minimo)
}

func suma(a, b int8) int8 {

	return a + b
}

func tipoEdad(edad uint8) string {
	var tipo string
	switch {
	case edad < 12:
		tipo = "Niñez"
	case edad < 18:
		tipo = "Adolescencia"
	default:
		tipo = "Madurez"
	}
	return tipo
}

//Firma de la funcion max y min
func maxymin(numeros []int8) (max int8, min int8) {

	for _, v := range numeros {
		if v > max {
			max = v
		} else {
			min = v
		}
	}

	return

}
