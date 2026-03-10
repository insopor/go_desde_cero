package modelos

import (
	"time"
)

/*
Estructuras:
Son parecidas al manejo de POO en otros lenguajes
pero en GO no existe como tal la POO.
Las estrcuturas y definiciones
siempre se definen a través de la palabra definida
type
*/

type User struct {
	Id            int
	Nombre        string
	FechaCreacion time.Time
	Estatus       bool
}

func (usuario User) AddUsuario(id int, nombre string, fechaCreacion time.Time, estatus bool) {
	usuario.Id = id
	usuario.Nombre = nombre
	usuario.FechaCreacion = fechaCreacion
	usuario.Estatus = estatus
}
