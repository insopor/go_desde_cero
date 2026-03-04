package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

// tablas de multiplicar
func Ejercicio02() {
	fmt.Println("ingrese el numero: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("es una cadena y no un numero, vuelva a intentarlo")
			Ejercicio02()
		} else {
			for i := 1; i < 11; i++ {
				fmt.Println(i, " X ", numero, " = ", i*numero)
			}
		}

	}
}

func Ejercicio02ptilotta() string {

	fmt.Println("ingrese el numero: ")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("es una cadena y no un numero, vuelva a intentarlo")
				continue
			} else {
				break
			}
		}
	}
	for i := 1; i < 11; i++ {
		//impresion normal con printf pára crear la cadena
		//fmt.Printf("%d X %d = %d \n", i, numero, i*numero)

		//guardar en una variable que genera una cadena para guardar en cadena para guarldarlo en un archivo
		texto += fmt.Sprintf("%d X %d = %d \n", numero, i, i*numero)

	}
	//fmt.Println(texto)
	return texto
}
