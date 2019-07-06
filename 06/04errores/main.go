package main

//Aprenderemos a manejar los errores personalizados
import (
	"errors"
	"fmt"
)

func main() {
	/*err := errors.New("Mi mensaje de error")
	//Printf me permite saber que tipo de variable es
	fmt.Printf("%T\n", err)*/

	r, err := division(100, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
		//el return permite que no se ejecute el codigo de abajo
	}
	fmt.Println(r)

}

func division(dividendo, divisor float64) (resultado float64, err error) {

	if divisor == 0 {
		err = errors.New("No se puede dividir por cero")
	} else {
		resultado = dividendo / divisor
	}

	return resultado, err
}
