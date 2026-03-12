package deferPanic

import (
	"fmt"
	"log"
)

/*
DEFER: determina las instrucciones finales de la ejecucion.
En el método VemosDifer() la salida es así:
	este es un primer mensajer
	este es el segundo mensaje
	este es el mensaje final

PANIC: determina cuando va a parar el código y a su vez, este mnos manda el error.
En la funcion EjemploPanic podemos ver que cuando el encuentra el valor mnos manda un error:
goroutine 1 [running]:
github.com/go_desde_cero/deferPanic.EjemploPanic(...)
        /go_desde_cero/deferPanic/deferPanicErrores.go:26
main.main()
        /go_desde_cero/main.go:107 +0xfe
exit status 2

RECOVER: puede determinar un manejo de errores antes de que panic pare los procesos, es decir,
panic para todo pero al hacer el recover con defer, se puede seguir el ciclo
para grabar todo los logs de los componentes que pudieron llegar a fallar,
aunque parece que podemos hacer lo mismo solo con defer

NOTA: no hay try - catch en go
*/

func VemosDifer() {
	fmt.Println("este es un primer mensajer")
	defer fmt.Println("este es el mensaje final")

	fmt.Println("este es el segundo mensaje")
}

func EjemploPanic() {
	a := 1
	if a == 1 {
		panic("se encontro el valor 1")
	}
}

func EjemploRecover() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error que genero panic \n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("se encontro el valor 1")
	}
}
