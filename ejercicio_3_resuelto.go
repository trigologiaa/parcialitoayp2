package main

import "fmt"

func Dibujar_TrianguloResuelto(n int) {
	if n == 0 {
		return
	}
	Dibujar_Triangulo(n - 1)
	for range n {
		fmt.Print("*")
	}
	fmt.Println()
}
