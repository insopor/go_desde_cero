package main

import (
	"fmt"

	"github.com/insopor/go_desde_cero/variables"
)

var r1 bool
var r2 string

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()

	r1, r2 := variables.ConvertirATexto(10)

	fmt.Println(r1, " - ", r2)
}
