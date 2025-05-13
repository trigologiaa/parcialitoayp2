package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotarResuelto3(t *testing.T) {
	list := NewList[int]()
	for i := 1; i <= 5; i++ {
		list.Append(i)
	}
	before := list.String()
	expectedBefore := "List: [1] ↔ [2] ↔ [3] ↔ [4] ↔ [5]"
	assert.Equal(t, expectedBefore, before, "La lista antes de rotar debe ser %s", expectedBefore)
	list.RotarResuelto(3)
	after := list.String()
	expectedAfter := "List: [4] ↔ [5] ↔ [1] ↔ [2] ↔ [3]"
	assert.Equal(t, expectedAfter, after, "La lista después de rotar debe ser %s", expectedAfter)
}

// func TestRotarResuelto3YDespuesMenos7(t *testing.T) {
// 	list := NewList[int]()
// 	for i := 1; i <= 5; i++ {
// 		list.Append(i)
// 	}
// 	initial := list.String()
// 	expectedInitial := "List: [1] ↔ [2] ↔ [3] ↔ [4] ↔ [5]"
// 	assert.Equal(t, expectedInitial, initial, "Lista inicial incorrecta")
// 	list.RotarResuelto(3)
// 	afterFirstRotation := list.String()
// 	expectedFirst := "List: [4] ↔ [5] ↔ [1] ↔ [2] ↔ [3]"
// 	assert.Equal(t, expectedFirst, afterFirstRotation, "Después de RotarResuelto(3) incorrecto")
// 	list.RotarResuelto(-7)
// 	final := list.String()
// 	expectedFinal := "List: [2] ↔ [3] ↔ [4] ↔ [5] ↔ [1]"
// 	assert.Equal(t, expectedFinal, final, "Después de RotarResuelto(-7) incorrecto")
// }

func TestRotar3(t *testing.T) {
	list := NewList[int]()
	for i := 1; i <= 5; i++ {
		list.Append(i)
	}
	before := list.String()
	expectedBefore := "List: [1] ↔ [2] ↔ [3] ↔ [4] ↔ [5]"
	assert.Equal(t, expectedBefore, before, "La lista antes de rotar debe ser %s", expectedBefore)
	list.Rotar(3)
	after := list.String()
	expectedAfter := "List: [4] ↔ [5] ↔ [1] ↔ [2] ↔ [3]"
	assert.Equal(t, expectedAfter, after, "La lista después de rotar debe ser %s", expectedAfter)
}

func TestRotar3YDespuesMenos7(t *testing.T) {
	list := NewList[int]()
	for i := 1; i <= 5; i++ {
		list.Append(i)
	}
	initial := list.String()
	expectedInitial := "List: [1] ↔ [2] ↔ [3] ↔ [4] ↔ [5]"
	assert.Equal(t, expectedInitial, initial, "Lista inicial incorrecta")
	list.Rotar(3)
	afterFirstRotation := list.String()
	expectedFirst := "List: [4] ↔ [5] ↔ [1] ↔ [2] ↔ [3]"
	assert.Equal(t, expectedFirst, afterFirstRotation, "Después de Rotar(3) incorrecto")
	list.Rotar(-7)
	final := list.String()
	expectedFinal := "List: [2] ↔ [3] ↔ [4] ↔ [5] ↔ [1]"
	assert.Equal(t, expectedFinal, final, "Después de Rotar(-7) incorrecto")
}
