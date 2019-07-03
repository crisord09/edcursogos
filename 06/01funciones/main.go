package main

import "fmt"

func main() {
	saludar("Christian", 40)
}

func saludar(nombre string, edad uint8) {

	if edad > 30 {
		fmt.Println("HOLA", nombre, "Tienes esta edad putito", edad)
	} else {
		fmt.Println("HOLA", nombre, "eres menor de edad trasero")
	}

}
