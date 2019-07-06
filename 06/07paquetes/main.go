package main

import (
	"fmt"

	"./despedida"
	"./saludar"
)

func main() {
	nombre := "Daniel"
	saludar.Saludar(nombre)

	saludar.MeVes = "Esto es un texto asignado desde el main"
	fmt.Println(saludar.MeVes)

	despedida.Despedirse(nombre)
}
