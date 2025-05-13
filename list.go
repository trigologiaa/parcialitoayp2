package main

import "fmt"

type List[T comparable] struct {
	centinelaCabeza *Node[T]
	centinelaCola   *Node[T]
	tamanio         int
}

// NewList crea una nueva lista vacía con nodos centinelas
func NewList[T comparable]() *List[T] {
	var zero T
	centinelaCabeza := NewNode(zero)
	centinelaCola := NewNode(zero)
	centinelaCabeza.SetNext(centinelaCola)
	centinelaCola.SetPrev(centinelaCabeza)
	return &List[T]{centinelaCabeza: centinelaCabeza, centinelaCola: centinelaCola}
}

func (l *List[T]) Head() *Node[T] {
	if l.IsEmpty() {
		return nil
	}
	return l.centinelaCabeza.Next()
}

func (l *List[T]) Tail() *Node[T] {
	if l.IsEmpty() {
		return nil
	}
	return l.centinelaCola.Prev()
}

func (l *List[T]) Size() int {
	return l.tamanio
}

func (l *List[T]) IsEmpty() bool {
	return l.tamanio == 0
}

func (l *List[T]) Clear() {
	l.centinelaCabeza.SetNext(l.centinelaCola)
	l.centinelaCola.SetPrev(l.centinelaCabeza)
	l.tamanio = 0
}

func (l *List[T]) Prepend(data T) {
	l.centinelaCabeza.InsertAfter(data)
	l.tamanio++
}

func (l *List[T]) Append(data T) {
	l.centinelaCola.InsertBefore(data)
	l.tamanio++
}

func (l *List[T]) RemoveFirst() {
	if l.IsEmpty() {
		return
	}
	node := l.centinelaCabeza.Next()
	node.Remove()
	l.tamanio--
}

func (l *List[T]) RemoveLast() {
	if l.IsEmpty() {
		return
	}
	node := l.centinelaCola.Prev()
	node.Remove()
	l.tamanio--
}

func (l *List[T]) Remove(data T) {
	for node := l.centinelaCabeza.Next(); node != l.centinelaCola; node = node.Next() {
		if node.Data() == data {
			node.Remove()
			l.tamanio--
			return
		}
	}
}

func (l *List[T]) Find(data T) *Node[T] {
	for node := l.centinelaCabeza.Next(); node != l.centinelaCola; node = node.Next() {
		if node.Data() == data {
			return node
		}
	}
	return nil
}

func (l *List[T]) String() string {
	if l.IsEmpty() {
		return "List: []"
	}
	result := "List: "
	for node := l.Head(); node != l.centinelaCola; node = node.Next() {
		result += fmt.Sprintf("[%v]", node.Data())
		if node.Next() != l.centinelaCola {
			result += " ↔ "
		}
	}
	return result
}
