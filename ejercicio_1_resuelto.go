package main

func (n *Node[T]) InsertAfterResuelto(dato T) {
	newNode := NewNode(dato)
	newNode.prev = n
	newNode.next = n.next
	if n.next != nil {
		n.next.prev = newNode
	}
	n.next = newNode
}

func (n *Node[T]) InsertBeforeResuelto(dato T) {
	newNode := NewNode(dato)
	newNode.next = n
	newNode.prev = n.prev
	if n.prev != nil {
		n.prev.next = newNode
	}
	n.prev = newNode
}
