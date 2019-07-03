package main

import "fmt"

func main() {

	//numeros := []string{"cero", "uno", "dos", "tres"}
	/*nombres := map[string]string{
		"es": "España",
		"co": "Colombia",
		"br": "Brazil",
	}

	for _, v := range nombres {
		fmt.Println(v)
	}*/

	/*for i := range numeros {
		fmt.Println(i)
	}*/

	frase := "Hola Mundo"

	for _, v := range frase {
		fmt.Println(string(v))
	}

	for _, entero := range []int{15, 13, 24, 85} {
		fmt.Println(entero)
	}

}
