package main

import (
	"fmt"
	//"runtime"

	//"github.com/insopor/go_desde_cero/archivos"
	//"github.com/insopor/go_desde_cero/ejercicios"

	//"github.com/insopor/go_desde_cero/teclado"
	//"github.com/insopor/go_desde_cero/arreglosslice"
	//"github.com/insopor/go_desde_cero/funciones"
	//"github.com/insopor/go_desde_cero/iteraciones"
	//"github.com/insopor/go_desde_cero/mapas"

	"github.com/insopor/go_desde_cero/deferPanic"
	ei "github.com/insopor/go_desde_cero/ejerciciosInterface"

	//gorutines "github.com/insopor/go_desde_cero/goRutines"
	"github.com/insopor/go_desde_cero/modelos"
	"github.com/insopor/go_desde_cero/usuarios"
	"github.com/insopor/go_desde_cero/variables"

	//"github.com/insopor/go_desde_cero/webserver"
	"github.com/insopor/go_desde_cero/middelware"
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
	/*
		sisOp := runtime.GOOS
		fmt.Println(sisOp)
		if sisOp == "darwin" || sisOp == "OS X" {
			fmt.Println("es mac")
		} else {
			fmt.Println("no es mac")
		}
	*/

	//forma de GO: se declara la variable dentro del if y despues se evalua dentro de la misma linea
	/*
		if sisOp := runtime.GOOS; sisOp == "Darwin." || sisOp == "OS X." {
			fmt.Println("esto no es windows")
		} else {
			fmt.Println("esto es windows")
		}
	*/

	// este es el rtipo switch que es para varios casos
	/*
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
	*/
	//fmt.Println(ejercicios.Ejercicio01("20"))

	//teclado.IngresoNumeros()

	//iteraciones.Iterar()
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
	//funciones.Exponencia(5)
	//arreglosslice.MuestraTabla()
	//arreglosslice.MuestraMatriz()
	//arreglosslice.MuestraSlice()
	//arreglosslice.SliceAPartirDeUnArreglo()
	//arreglosslice.Capacidad()
	//mapas.MostrarMapas()

	usuarios.AltaUsuario()

	Pedro := new(modelos.Hombre)
	ei.HumanosRespirando(Pedro)

	//ejerciciosinterface.HumanosRespirando()
	Maria := new(modelos.Mujer)
	ei.HumanosRespirando(Maria)

	deferPanic.VemosDifer()
	//deferPanic.EjemploPanic()
	//deferPanic.EjemploRecover()

	//inicia la go rutine
	//go gorutines.MiNombreLento("arturo")

	//la ejecución del programa no termina
	/*
		fmt.Println("estoy aqui")
		var x string
		fmt.Scanln(&x)
	*/

	/*
		//de esta manera podemos hacer que la ejecución termine hasta que termine de ejecutarse la rutina
		canal1 := make(chan bool) //creamos el canal para inyectar
		go gorutines.MiNombreLentoChanel("arturo", canal1)
		fmt.Println("estoy aqui")
		<-canal1 // asi se declara el await o la espera a que se termine la rutina

		//también de esta manera podemos declarar el await dentro de una funcion anonima
		// y dentro de la funcion también podemos declarar código extra
		canal2 := make(chan bool)
		go gorutines.MiNombreLentoChanel("beto", canal2)
		defer func() {
			<-canal2
		}()
		fmt.Println("estoy aqui 2")
	*/

	//webserver.MiWebSerbver()

	middelware.MiMiddleware()

}
