package usuarios

import (
	"fmt"
	"time"

	"github.com/insopor/go_desde_cero/modelos"
)

func AltaUsuario() {
	usuario := new(modelos.User)
	usuario.AddUsuario(10, "pablo", time.Now(), true)
	fmt.Println(usuario)
}
