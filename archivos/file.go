package archivos

import (
	"bufio"
	"fmt"
	"os"

	"github.com/insopor/go_desde_cero/ejercicios"
)

func GrabarTabla() {
	var texto string = ejercicios.Ejercicio02ptilotta()
	archivo, err := os.Create("./archivos/txt/tabla.txt")
	if err != nil {
		fmt.Println("error al crear el archivo " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

var nombreArchivo string = "./archivos/txt/tabla.txt"

func Sumatabla() {

	var texto string = ejercicios.Ejercicio02ptilotta()
	if Appendo(nombreArchivo, texto) == false {
		fmt.Println("error al concatener contenido")
	}

}

func Appendo(archivo string, texto string) bool {

	arch, err := os.OpenFile(nombreArchivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("error en el append " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("error durante el writestring " + err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		fmt.Println("error en la lectura del archivo " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
}
