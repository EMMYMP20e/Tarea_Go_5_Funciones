package main

import "fmt"

func fibonacci(x, y, limite, actual uint) uint {
	if actual >= limite {
		return y
	}
	suma := x + y
	x = y
	y = suma
	actual++
	return fibonacci(x, y, limite, actual)
}

func inicio(limite uint) uint {
	return fibonacci(0, 1, limite, 1)
}

func main() {
	posicion := uint(10)
	fmt.Println("El Número ", posicion, " de la serie de fibonacci es: ", inicio(posicion))
}
