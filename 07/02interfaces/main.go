// las interfaces son un contrato que se puede implementar en cualquier tipo de dato en GO
// Nos habla un contrato que existe que se puede implementar
package main

import "./animales"

func main() {

	var p animales.Perro
	var g animales.Gato

	p.Nombre = "Firulais"
	g.Nombre = "Manchas"

	//AdoptarPerro(p)
	//AdoptarGato(g)
	AdoptarMascota(p)
	AdoptarMascota(g)

}

/*func AdoptarPerro(p animales.Perro) {
	p.Comunicarse()
}

func AdoptarGato(g animales.Gato) {
	g.Comunicarse()
}*/

func AdoptarMascota(m animales.Mascota) {
	m.Comunicarse()
}
