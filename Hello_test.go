package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	res := "Hello World from Go !!"
	if got != "Hello World from Go !!" {
		t.Errorf("HelloWorld() = %s; want %s", got, res)
	}

}
