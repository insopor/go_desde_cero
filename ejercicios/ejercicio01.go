package ejercicios

import (
	"strconv"
)

func Ejercicio01(valor string) (int, string) {
	var mensaje string
	numerico, error := strconv.Atoi(valor)
	//null no existe en go, solo sirve nil
	if error != nil {
		mensaje = "no se puede convertir"
		return 1, mensaje
	} else {
		if numerico > 100 {
			mensaje = "el valor es mayor a 100"
			return numerico, mensaje
		} else {
			mensaje = "el numero es menor a 100"
			return numerico, mensaje
		}

	}

}
