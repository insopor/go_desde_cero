package funciones

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Recursion1() {
	var numero int
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("es una cadena y no un numero, vuelva a intentarlo")
			Recursion1()
		} else {
			for i := 1; i < 11; i++ {
				fmt.Println(i, " X ", numero, " = ", i*numero)
			}
		}

	}
}

func Exponencia(valor int) {
	if valor > 100 {
		return
	}

	fmt.Println(valor)
	Exponencia(valor * 2)
}
