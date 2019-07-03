package main

import "fmt"

func main() {

	/*
		idiomas := make(map[string]string)
		idiomas["es"] = "Español"
		idiomas["en"] = "Ingles"
		idiomas["fr"] = "Frances"
	*/

	idiomas := map[string]string{
		"es": "Español",
		"en": "ingles",
		"fr": "Frances",
		"pt": "Portugues",
	}
	/*
		fmt.Println(idiomas["fr"])

		numeros := map[int]string{
			1:    "Uno es un numero chiquito",
			2016: "Esto es otro número",
			-4:   "Es un número negativo",
		}

		fmt.Println(numeros[1])
	*/
	if idioma, ok := idiomas["pt"]; ok {
		fmt.Println("La posición pt si existe y vale", idioma)
	} else {
		fmt.Println("La posición pt No existe")
	}
}
