package gorutines

import (
	"fmt"
	"strings"
	"time"
)

/*
GO RUTINES: son rutinas que se ejecutan mientras otras funciones tambien se ejecutan a la vez,
parecido a los hilos en Java. Ladiferencia es que si la ejecución termina antes
que la funcion definida como "go rutine", entonces la go rutine no termina su proceso

*/

// funcion que hace que el nombre se imprima letra por letra de manera lenta
func MiNombreLento(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}

func MiNombreLentoChanel(nombre string, canalUno chan bool) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}

	canalUno <- true
}
