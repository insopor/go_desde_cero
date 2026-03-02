package variables

import (
	"fmt"
	"time"
)

/*
//cuando una funcion es declarada con mayusocula al inicio, significa que será publica
//las funciones publicas podrán ser vistas por cualquier paquete que invoque este paquete
*/

var Nombre string
var Estado bool
var Sueldo float32
var fecha time.Time

func RestoVariables() {
	Nombre = "pedro"
	Estado = true
	Sueldo = 1500.10
	fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(fecha)

}
