package main

func (n *Node[T]) InsertAfter(dato T) {
	nuevo := &Node[T]{data: dato, next: n.next, prev: n}
	if n.next != nil {
		n.next.prev = nuevo
	}
	n.next = nuevo
}

func (n *Node[T]) InsertBefore(dato T) {
	nuevo := &Node[T]{data: dato, next: n, prev: n.prev}
	if n.prev != nil {
		n.prev.next = nuevo
	}
	n.prev = nuevo
}
