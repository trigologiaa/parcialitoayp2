package main

import "fmt"

func Dibujar_Triangulo(n int) {
	dibujarLinea(1, n)
}

func dibujarLinea(actual, total int) {
	if actual > total {
		return
	}
	for i := 0; i < actual; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	dibujarLinea(actual+1, total)
}
