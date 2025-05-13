package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertAfterResuelto(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node1.SetNext(node2)
	node2.SetPrev(node1)
	node1.InsertAfterResuelto(3)
	node3 := node1.Next()
	assert.Equal(t, 3, node3.Data(), "El nodo insertado debe tener valor 3")
	assert.Equal(t, node1, node3.Prev(), "El nodo nuevo debe tener como prev a node1")
	assert.Equal(t, node2, node3.Next(), "El nodo nuevo debe tener como next a node2")
	assert.Equal(t, node3, node2.Prev(), "node2 debe tener como prev al nodo nuevo")
	assert.Equal(t, node3, node1.Next(), "node1 debe tener como next al nodo nuevo")
}

func TestInsertBeforeResuelto(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node1.SetNext(node2)
	node2.SetPrev(node1)
	node2.InsertBeforeResuelto(3)
	node3 := node2.Prev()
	assert.Equal(t, 3, node3.Data(), "El nodo insertado debe tener valor 3")
	assert.Equal(t, node1, node3.Prev(), "El nodo nuevo debe tener como prev a node1")
	assert.Equal(t, node2, node3.Next(), "El nodo nuevo debe tener como next a node2")
	assert.Equal(t, node3, node1.Next(), "node1 debe tener como next al nodo nuevo")
	assert.Equal(t, node3, node2.Prev(), "node2 debe tener como prev al nodo nuevo")
}

func TestInsertAfter(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node1.SetNext(node2)
	node2.SetPrev(node1)
	node1.InsertAfter(3)
	node3 := node1.Next()
	assert.Equal(t, 3, node3.Data(), "El nodo insertado debe tener valor 3")
	assert.Equal(t, node1, node3.Prev(), "El nodo nuevo debe tener como prev a node1")
	assert.Equal(t, node2, node3.Next(), "El nodo nuevo debe tener como next a node2")
	assert.Equal(t, node3, node2.Prev(), "node2 debe tener como prev al nodo nuevo")
	assert.Equal(t, node3, node1.Next(), "node1 debe tener como next al nodo nuevo")
}

func TestInsertBefore(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node1.SetNext(node2)
	node2.SetPrev(node1)
	node2.InsertBefore(3)
	node3 := node2.Prev()
	assert.Equal(t, 3, node3.Data(), "El nodo insertado debe tener valor 3")
	assert.Equal(t, node1, node3.Prev(), "El nodo nuevo debe tener como prev a node1")
	assert.Equal(t, node2, node3.Next(), "El nodo nuevo debe tener como next a node2")
	assert.Equal(t, node3, node1.Next(), "node1 debe tener como next al nodo nuevo")
	assert.Equal(t, node3, node2.Prev(), "node2 debe tener como prev al nodo nuevo")
}
