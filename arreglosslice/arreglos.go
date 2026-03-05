package arreglosslice

import (
	"fmt"
)

/*
En GO los arreglos son como en cualquier otro lenguaje
con una longitud definida de principioa fin
*/

// vectort o arreglo normal
var tabla [10]int = [10]int{20, 21, 22}

// matriz o arreglo multiple multiple
var matriz [10][10]int

func MuestraTabla() {
	for i := 0; i < len(tabla); i++ {
		tabla[i] = i
	}

	tabla[3] = 33
	tabla[4] = 54

	tabla2 := [10]string{"a", "b", "c"}

	fmt.Println(tabla)
	fmt.Println(tabla2)
}

func MuestraMatriz() {
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz); j++ {
			matriz[i][j] = i
			fmt.Println(matriz[i][j])
			//fmt.Println(matriz[i])
			//fmt.Println(matriz[j])
		}
		fmt.Println()
	}
}
