package mapas

import (
	"fmt"
)

func MostrarMapas() {
	fmt.Println()
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["mexico"] = "df"
	paises["argentina"] = "buenos aires"
	paises["eua"] = "washington"

	fmt.Println(paises)           //se imprime el mapa completo
	fmt.Println(paises["mexico"]) //se imprime el valor de la clave

	campeonato := map[string]int{
		"barcelona":   20,
		"real madrid": 21,
		"chivas":      30,
		"boca":        28,
	}

	//NOTA: los mapas se recorren de manera asincrona asi que no sale ordenado de niunguna manera
	//se muestra el mapa completo de campeonato
	fmt.Println(campeonato)

	//se recorre el mapa de campeonato
	for equipo, puntaje := range campeonato {
		fmt.Printf("equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	//quitar un elemento de un mapa
	delete(campeonato, "real madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["juventus"]
	fmt.Printf("el porcentaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)
}
