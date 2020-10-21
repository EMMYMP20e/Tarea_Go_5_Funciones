package main

import "fmt"

func intercambia(x, y *int) {
	a := *x
	*x = *y
	*y = a
}

func main() {
	a := 10
	b := 5
	fmt.Println("a: ", a, " |b: ", b)
	intercambia(&a, &b)
	fmt.Println("Intercambio: ")
	fmt.Println("a: ", a, " |b: ", b)
}
