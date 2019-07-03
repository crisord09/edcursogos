package main

import "fmt"

//Persona es uun tipo de dato estructural
type Persona struct {
	//Similar a las clases de otros legunajes de programación
	Nombre string
	Edad   uint8
	Emails []string
}

func main() {
	/*
		var persona1 Persona
		persona1.Nombre = "Pedro"
		persona1.Edad = 20
	*/

	emailss := []string{"correo1", "correo2"}

	persona2 := Persona{
		Nombre: "Pablo",
		Edad:   31,
		Emails: emailss,
	}

	fmt.Print(persona2)
}
