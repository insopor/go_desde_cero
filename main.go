package main

import (
	"fmt"
	"runtime"

	//"github.com/insopor/go_desde_cero/archivos"
	"github.com/insopor/go_desde_cero/ejercicios"

	//"github.com/insopor/go_desde_cero/teclado"
	"github.com/insopor/go_desde_cero/arreglosslice"
	"github.com/insopor/go_desde_cero/funciones"
	"github.com/insopor/go_desde_cero/iteraciones"
	"github.com/insopor/go_desde_cero/mapas"
	"github.com/insopor/go_desde_cero/variables"
)

var r1 bool
var r2 string

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()

	r1, r2 := variables.ConvertirATexto(10)
	fmt.Println(r1, " - ", r2)

	//forma tradicional:
	// se declara la variable y despues se evalua el resultado
	sisOp := runtime.GOOS
	fmt.Println(sisOp)
	if sisOp == "darwin" || sisOp == "OS X" {
		fmt.Println("es mac")
	} else {
		fmt.Println("no es mac")
	}

	//forma de GO: se declara la variable dentro del if y despues se evalua dentro de la misma linea
	/*
		if sisOp := runtime.GOOS; sisOp == "Darwin." || sisOp == "OS X." {
			fmt.Println("esto no es windows")
		} else {
			fmt.Println("esto es windows")
		}
	*/

	// este es el rtipo switch que es para varios casos
	switch sisOp := runtime.GOOS; sisOp {
	case "windows":
		fmt.Println("es windows")
	case "darwi":
		fmt.Println("es darwin")
	case "OSX":
		fmt.Println("es mac os")
	default:
		fmt.Printf("%s \n ", sisOp)
	}

	fmt.Println(ejercicios.Ejercicio01("20"))

	//teclado.IngresoNumeros()

	iteraciones.Iterar()
	//iteraciones.IterarDos()
	//iteraciones.IterarTres()
	//iteraciones.IterarBreak()
	//iteraciones.IterarContinue()

	//ejercicios.Ejercicio02()
	//fmt.Println(ejercicios.Ejercicio02ptilotta())

	//este metodo suma tabla es el que crea la escritura sobre el archivo
	//archivos.Sumatabla()

	//este método de leoarchivo es para leer el archivo que previamente se hizo con sumatabla
	//archivos.LeoArchivo()

	//funciones.Calculos()
	//funciones.LlamarClousure()
	//funciones.Recursion1()
	funciones.Exponencia(5)
	arreglosslice.MuestraTabla()
	//arreglosslice.MuestraMatriz()
	arreglosslice.MuestraSlice()
	arreglosslice.SliceAPartirDeUnArreglo()
	arreglosslice.Capacidad()
	mapas.MostrarMapas()
}
