package main

import (
	"log"
	"net/http"
)

func main() {
	/*
		http.Handle("/", http.FileServer(http.Dir("public")))
		log.Println("Ejecutando server en http://localhost:8080")
		log.Println(http.ListenAndServe(":8080", nil))
	*/
	//creamos un servidor para la eficiencia con un multiplexor
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	log.Println("Ejecutando server en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))

}
