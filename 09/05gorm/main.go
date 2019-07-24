package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//Producto contiene los datos de un aritculo
type Producto struct {
	gorm.Model   //ID, CreatedAt, UpdatedAt, DeletedAt
	CodigoBarras string
	Precio       uint
}

func main() {
	db, err := gorm.Open("mysql", "root:752mile@/edcurso?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("Error en la conexión a la base de datos")
	}
	defer db.Close()
	fmt.Println("Se conectó a la base de datos")
	/*
		db.CreateTable(&Producto{})
		fmt.Println("Se creo la tabla productos en la base de datos")
	*/
	/*
		db.Create(&Producto{
			CodigoBarras: "otroCodigo",
			Precio:       4500,
		})
	*/

	var p Producto
	db.First(&p, 2)
	fmt.Println("Se extrayo la informacion del primer registro", p)
	fmt.Println("Precio del prodcuto solicitado", p.Precio)

	p.Precio = 5000
	db.Save(&p)
	fmt.Println("El nuevo precio del producto es: ", p.Precio)

}
