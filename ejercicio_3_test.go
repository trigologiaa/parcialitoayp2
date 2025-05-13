package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestDibujarTrianguloResuelto(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, ""},
		{1, "*\n"},
		{2, "*\n**\n"},
		{3, "*\n**\n***\n"},
	}
	for _, tt := range tests {
		output := captureOutput(func() {
			Dibujar_TrianguloResuelto(tt.input)
		})
		assert.Equal(t, tt.expected, output, "Failed on input %d", tt.input)
	}
}

func TestDibujarTriangulo(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, ""},
		{1, "*\n"},
		{2, "*\n**\n"},
		{3, "*\n**\n***\n"},
	}
	for _, tt := range tests {
		output := captureOutput(func() {
			Dibujar_Triangulo(tt.input)
		})
		assert.Equal(t, tt.expected, output, "Failed on input %d", tt.input)
	}
}
