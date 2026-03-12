package middelware

import (
	"fmt"
)

func sumarMiddle(a, b int) int {
	return a + b
}

func restarMiddle(a, b int) int {
	return a - b
}

func multiMiddle(a, b int) int {
	return a * b
}

func divMiddle(a, b int) int {
	return a / b
}

func MiMiddleware() {
	fmt.Println("inicio")
	result := operacionesMidd(sumarMiddle)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restarMiddle)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(multiMiddle)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(divMiddle)(2, 3)
	fmt.Println(result)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("inicio de operacion")
		return f(a, b)
	}
}
