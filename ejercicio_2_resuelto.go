package main

func (l *List[T]) RotarResuelto(posiciones int) {
	rotarIzquierda := false
	if posiciones < 0 {
		posiciones = -posiciones
		rotarIzquierda = true
	}
	posiciones = posiciones % l.Size()
	if l.IsEmpty() || posiciones == 0 || l.Size() == 1 {
		return
	}
	if rotarIzquierda {
		for range posiciones {
			ultimo := l.Tail().Data()
			l.centinelaCabeza.next.InsertAfter(ultimo)
			l.tamanio++
			l.RemoveLast()
		}
	} else {
		for range posiciones {
			primero := l.Head().Data()
			l.centinelaCola.InsertBefore(primero)
			l.tamanio++
			l.RemoveFirst()
		}
	}
}
