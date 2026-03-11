package modelos

type Mujer struct {
	Hombre //asi se heredan las propiedades
}

func (m *Mujer) Respirar()      { m.respirando = true }
func (m *Mujer) Comer()         { m.comiendo = true }
func (m *Mujer) Pensar()        { m.pensando = true }
func (m *Mujer) Sexo() string   { return "mujer" }
func (m *Mujer) Estavivo() bool { return true }
