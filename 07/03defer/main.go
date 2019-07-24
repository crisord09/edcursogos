package main

import "fmt"

func main() {
	fmt.Println("Conectando a la base de datos...")
	//Es como una sala de espera donde todas las funciones estan a la espera de ejecutarse

	//Operaciones con la conexión
	defer cerrarBD()

	//Obliga a que si o si se ejecute la función
	fmt.Println("Consultamos información, set de datos")
	defer cesarrSetDeDatos()

	//En GO una función se termina cuando llega a la sentencia return
}

func cerrarBD() {
	fmt.Println("Cerrar la base de datos")
}

func cesarrSetDeDatos() {
	fmt.Println("Cerrar el set de datos")

}
