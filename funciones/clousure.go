package funciones

import (
	"fmt"
)

// tabla recibe un valor entero y explusa una funcion que regresa un entero
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func LlamarClousure() {
	tabladel := 2
	Mitabla := tabla(tabladel)
	for i := 1; i <= 10; i++ {
		fmt.Println(Mitabla())
	}
}
