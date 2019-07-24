// Handler func son un tipo de dato que permite converitr funciones en manejadores HTTP
//la firma debe de tener nombre la funcion recive el responwriter y el request, solamente creamos la funció
package main

import (
	"fmt"
	"log"
	"net/http"
)

//Es como nuestro controlador
func messageHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola este es un handlerfunc</h1>")
}

func messageHFCustom(message string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, message)
		},
	)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", messageHandlerFunc)
	//EL mensaje personalizado no tiene la firma que devuelve el responsewriter y request por eso solo utilizamos el Handle
	mux.Handle("/saludar", messageHFCustom("<h1>Hola amigos</h1>"))
	mux.Handle("/despedirse", messageHFCustom("<h1>Chau Amigos</h1>"))

	log.Println("Ejecutando server en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
