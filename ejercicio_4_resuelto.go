package main

func Maximo(arr []int, inicio, fin int) int {
	if inicio == fin {
		return arr[inicio]
	}
	mid := (inicio + fin) / 2
	max := Maximo(arr, inicio, mid)
	for i := mid + 1; i <= fin; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// A = 1, B = 2, C = 1 -> log b (a) < C -> O(n^c) = O(n)
