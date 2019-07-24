//El paralelismo consiste en utilizar todas las GOrutinas y obliga a la función main que ejecute todas las Gorutinas que yo haya determinado
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//Me ayuda a designar cuantos procesadores usare
	runtime.GOMAXPROCS(1)

	//no termine hasta que termine todas mis operaciones con gorutinas
	//Grupo de espera
	var wg sync.WaitGroup

	//

	numbers := []uint32{125424, 4, 235, 12341, 5321, 43211, 1241877, 32135647, 6544774, 41748}

	//Con esto le decimos cuantas goRutinas o(actividades) vamos a lanzzar
	wg.Add(len(numbers))

	fmt.Println("Comienza el proceso....")
	for _, v := range numbers {
		go func(a uint32) {
			//Con el defer le decimos que se ejecute si o si y con el Done que lo quite el proceso de el grupo de espera
			defer wg.Done()
			fmt.Println(a, EsPrimo(v))
		}(v)
	}
	//Debes de esperar a que termine todos los procesos que estan antes de la linea wait para pasar a la siguiente linea de código
	wg.Wait()
	fmt.Println("Termino el proceso")

}

func EsPrimo(a uint32) bool {
	c := 0
	var i uint32
	for i = 1; i <= a; i++ {
		if a%i == 0 {
			c++
		}
	}
	if c == 2 {
		return true
	}
	return false
}
