package main

func (l *List[T]) Rotar(k int) {
	if l.IsEmpty() || k == 0 || l.Size() == 1 {
		return
	}
	rotarIzquierda := false
	if k < 0 {
		k = -k
		rotarIzquierda = true
	}
	k = k % l.Size()
	if k == 0 {
		return
	}
	if rotarIzquierda {
		for range k {
			ultimo := l.Tail().Data()
			l.Head().InsertBefore(ultimo)
			l.tamanio++
			l.RemoveLast()
		}
	} else {
		for range k {
			primero := l.Head().Data()
			l.centinelaCola.InsertBefore(primero)
			l.tamanio++
			l.RemoveFirst()
		}
	}
}
