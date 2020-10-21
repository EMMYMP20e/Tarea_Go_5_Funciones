package main

import "fmt"

func sum(args ...int) int {
	mayor := args[0]
	for _, v := range args { // iterar sobre los parámetros como un slice
		if v > mayor {
			mayor = v
		}
	}
	return mayor
}

func main() {
	a := []int{4, 1, 0, -1, -2, 3, 8, 7, 2, 9, 6, 5}
	fmt.Println(sum(a...))
}
