// los canales perminten la comunicación entre Gorutinas (Permite que entre ellas se sincronicen o intercambien información para cumplir con la tarea)
// Comparte información con las actividades
package main

import "fmt"

func main() {
	bodegaOrigen := []string{"php", "javascript", "html", "css", "java", "bases de datos", "git"}
	bodegaDestino := []string{}

	//make nos ayuda a crar un canal, creamos con chan y el tipoi de dato

	fotocopiadora := make(chan string)
	fotocopiado := make(chan string)
	//los canales perminten comunicarse entre si a las Gorutinas de actividades

	go func() {
		for _, libro := range bodegaOrigen {
			fotocopiadora <- libro
		}
	}()

	//Se encarga de llevarlo a la fotocopiadora y sacar 1 copia del libro
	go func() {
		for {
			//Se entrega el contenido de la fotocopiadora a la variable libro
			libro := <-fotocopiadora

			//le vamos a enviar el contenido que tienen la variable libro
			fotocopiado <- libro

			//Agregar la nueva copia del libro a la bodega destino
			bodegaDestino = append(bodegaDestino, libro)
			if len(bodegaOrigen) == len(bodegaDestino) {
				// cerrar el canal fotocopiado
				close(fotocopiado)
			}

		}
	}()

	fmt.Println("*** Listado de libros fotocopiados ***")
	for {
		if libro, ok := <-fotocopiado; ok {
			fmt.Println(libro)
		} else {
			break
		}
	}

	// Se encarga de guardarlo en la bodega destino

}
