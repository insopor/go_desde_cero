package ejerciciosinterface

import (
	"github.com/insopor/go_desde_cero/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	hu.Sexo()
	//fmt.Println("soy " + hu.Sexo() + "y estoy respirando")
}
