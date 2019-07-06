//Funciones con minuscula nos son privadas con Mayuscula son publicas

package saludar

import "fmt"

//MeVes es para utilizar desde otro paquete
var MeVes string

//noMeVes no puedes tulizar desde otro paquete
var noMeVes string

//Saluda a una persona
func Saludar(nombre string) {
	fmt.Println("Hola", nombre)
}
