package iteraciones

import (
	"fmt"
)

func Iterar() {
	fmt.Println("Itera normal")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}

func IterarDos() {
	fmt.Println("Itera de 5 en 5")
	for i := 0; i < 15; i += 5 {
		fmt.Println(i)
	}
}

func IterarTres() {
	fmt.Println("Itera al reves")
	for i := 100; i > 70; i -= 5 {
		fmt.Println(i)
	}
}

func IterarBreak() {
	fmt.Println("Itera y corta en 6")
	for i := 10; i > 1; i-- {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
}

func IterarContinue() {
	fmt.Println("Itera y salta el 6")
	for i := 10; i > 1; i-- {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
}
