package main

import "testing"

func TestMaximo(t *testing.T) {
	casos := []struct {
		nombre   string
		arr      []int
		inicio   int
		fin      int
		esperado int
	}{
		{"Caso 1: array ordenado", []int{1, 2, 3, 4, 5}, 0, 4, 5},
		{"Caso 2: array desordenado", []int{3, 7, 2, 9, 1}, 0, 4, 9},
		{"Caso 3: un solo elemento", []int{42}, 0, 0, 42},
		{"Caso 4: negativos", []int{-10, -20, -30}, 0, 2, -10},
		{"Caso 5: subarreglo", []int{5, 1, 9, 3, 7}, 1, 3, 9},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			resultado := Maximo(c.arr, c.inicio, c.fin)
			if resultado != c.esperado {
				t.Errorf("Esperado %d, pero se obtuvo %d", c.esperado, resultado)
			}
		})
	}
}
