package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	fmt.Println("ingrese numero 1: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("el dato ingresado es incorrecto " + err.Error())
		}

	}

	fmt.Println("ingrese numero 2: ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("el dato ingresado es incorrecto " + err.Error())
		}

	}

	fmt.Println("ingrese leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)
}
