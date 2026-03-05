package arreglosslice

import (
	"fmt"
)

/*
En GO los Slices son como las Listas en Java,
arreglos que se ajustan o mejor dicho, que son dinamicos
*/

//la declaracion de un slice es como iun arreglo
//pero sin el size definido como en los arreglos normales

/*
dimension y capacidad
la capacidad es el espacio de memoria que go reserva
es decir, nosotros no solo podemos decir a go
que cuando compile, solo compile la memoria que reserve
con el arreglo si no que podemos declarar que
vamos a ocupar 5 elemntos pero el arreglo puede tener 15 elemntos mas
en un ambiiente dinamico
*/

var sliceTabla []int = []int{2, 5, 4}

func MuestraSlice() {
	fmt.Println(sliceTabla)
}

var arregloALlenar [10]int

func SliceAPartirDeUnArreglo() {
	for i := 0; i < len(arregloALlenar); i++ {
		arregloALlenar[i] = i
	}

	porcion1 := arregloALlenar[3:]  // el slice se crea desde el arreglo desde la posicion 3
	porcion2 := arregloALlenar[:5]  // el slice se crea desde el arreglo desde la posicion 0 a la 5
	porcion3 := arregloALlenar[3:8] // el slice se crea desde el arreglo desde la posicion 3 hasta la 8

	fmt.Println(porcion1)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	//en elementos damos a entender con make la capacidad de uso y la disponibilidad de memoria
	elementos := make([]int, 5, 20)
	fmt.Printf("largo %d, capacidad %d", len(elementos), cap(elementos))
	fmt.Println()
	//en nums hacemos lo mismo que en elementos pero llenamos el slice
	//si el slice es de 100 llenado dinamicamente, go ocupara 128 ya que
	// nosotros no reservamos la memoria sino que go ya la tiene previamente hecha
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("largo %d, capacidad %d", len(nums), cap(nums))

}
