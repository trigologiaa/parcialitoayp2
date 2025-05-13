package main

// Node representa un nodo de una lista enlazada doble.
type Node[T comparable] struct {
	data T
	next *Node[T]
	prev *Node[T]
}

// NewDoubleLinkedNode crea un nuevo nodo de lista enlazada doble con el dato especificado.
func NewNode[T comparable](data T) *Node[T] {
	return &Node[T]{data: data}
}

// SetData establece el dato almacenado en el nodo.
func (n *Node[T]) SetData(data T) {
	n.data = data
}

// Data devuelve el dato almacenado en el nodo.
func (n *Node[T]) Data() T {
	return n.data
}

// SetNext establece el nodo siguiente al nodo actual.
func (n *Node[T]) SetNext(next *Node[T]) {
	n.next = next
}

// Next devuelve el nodo siguiente al nodo actual.
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// HasNext devuelve true si el nodo actual tiene asignado un nodo siguiente.
func (n *Node[T]) HasNext() bool {
	return n.next != nil
}

// SetPrev establece el nodo anterior al nodo actual.
func (n *Node[T]) SetPrev(newPrev *Node[T]) {
	n.prev = newPrev
}

// Prev devuelve el nodo anterior al nodo actual.
func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

// HasPrev devuelve true si el nodo actual tiene asignado un nodo previo.
func (n *Node[T]) HasPrev() bool {
	return n.prev != nil
}

// Remove elimina el nodo actual de la lista.
func (n *Node[T]) Remove() {
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
	n.prev = nil
	n.next = nil
}

// RemoveNext elimina el nodo siguiente al nodo actual.
func (n *Node[T]) RemoveNext() {
	if n.next != nil {
		nextNext := n.next.next
		if nextNext != nil {
			nextNext.prev = n
		}
		n.next = nextNext
	}
}

// RemoveLast elimina el nodo anterior al nodo actual.
func (n *Node[T]) RemoveLast() {
	if n.prev != nil {
		prevPrev := n.prev.prev
		if prevPrev != nil {
			prevPrev.next = n
		}
		n.prev = prevPrev
	}
}
