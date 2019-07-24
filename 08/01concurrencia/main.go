package main

import (
	"fmt"
	"time"
)

func main() {
	var h string
	//Ejecute mostrar numerso como uyna gorutina como una manera independiente
	//Gorutina Permite ejecutar un proceso de manera independiente (eso es la concurrencia en GO)
	go MostrarNumeros()
	fmt.Print("Digita algo ")
	fmt.Scan(&h)
	fmt.Println("Digitaste", h)

}

func MostrarNumeros() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
