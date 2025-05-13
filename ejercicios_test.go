package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	list := NewList[int]()
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestListPrepend(t *testing.T) {
	list := NewList[string]()
	list.Prepend("1")
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
	list.Prepend("2")
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "2", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
}

func TestListAppend(t *testing.T) {
	list := NewList[string]()
	list.Append("1")
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
	list.Append("2")
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "2", list.Tail().Data())
}

func TestListClear(t *testing.T) {
	list := NewList[int]()
	list.Append(1)
	list.Append(2)
	list.Clear()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestListIsEmpty(t *testing.T) {
	list := NewList[int]()
	assert.True(t, list.IsEmpty())
	list.Append(1)
	assert.False(t, list.IsEmpty())
}

func TestListSize(t *testing.T) {
	list := NewList[int]()
	assert.Equal(t, 0, list.Size())
	list.Append(1)
	assert.Equal(t, 1, list.Size())
	list.Append(2)
	assert.Equal(t, 2, list.Size())
}

func TestListHead(t *testing.T) {
	list := NewList[int]()
	assert.Nil(t, list.Head())
	list.Append(1)
	assert.Equal(t, 1, list.Head().Data())
	list.Append(2)
	assert.Equal(t, 1, list.Head().Data())
}

func TestListTail(t *testing.T) {
	list := NewList[int]()
	assert.Nil(t, list.Tail())
	list.Append(1)
	assert.Equal(t, 1, list.Tail().Data())
	list.Append(2)
	assert.Equal(t, 2, list.Tail().Data())
}

func TestListRemove(t *testing.T) {
	list := NewList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Remove(2)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())
	list.Remove(1)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 3, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())
	list.Remove(3)
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestListRemoveLastElement(t *testing.T) {
	list := NewList[int]()
	list.Append(1)
	list.Append(2)
	list.Remove(2)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 1, list.Tail().Data())
}

func TestListRemoveFirst(t *testing.T) {
	list := NewList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.RemoveFirst()
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 2, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())
	list.RemoveFirst()
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 3, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())
	list.RemoveFirst()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestListRemoveLast(t *testing.T) {
	list := NewList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.RemoveLast()
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 2, list.Tail().Data())
	list.RemoveLast()
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 1, list.Tail().Data())
	list.RemoveLast()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestListAllRemovesOnEmptyList(t *testing.T) {
	list := NewList[int]()
	list.RemoveFirst()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
	list.RemoveLast()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
	list.Remove(1)
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestListFind(t *testing.T) {
	list := NewList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	nodo := list.Find(2)
	assert.NotNil(t, nodo)
	assert.Equal(t, 2, nodo.Data())
	nodo = list.Find(4)
	assert.Nil(t, nodo)
}

func TestListStringOnEmpty(t *testing.T) {
	list := NewList[int]()
	assert.Equal(t, "List: []", list.String())
}

func TestListString(t *testing.T) {
	list := NewList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	assert.Equal(t, "List: [1] ↔ [2] ↔ [3]", list.String())
}
