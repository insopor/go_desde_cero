package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var leyenda string
var err error

func Ejercicio02() {
	fmt.Println("ingrese el numero: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("el dato ingresado no es correcto " + err.Error())
		} else {
			for i := 1; i < 11; i++ {
				fmt.Println(i, "X", numero, i*numero)
			}
		}

	}
}
