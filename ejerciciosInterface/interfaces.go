package ejerciciosinterface

import (
	"fmt"
	//"github.com/insopor/go_desde_cero/interfaces"
	"github.com/insopor/go_desde_cero/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Println("soy " + hu.Sexo() + "y estoy respirando")
}
