package main

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestHelloWorld(t *testing.T) {
	g := Goblin(t)
	g.Describe("Activity 1", func() {
		g.Assert(HelloWorld()).Equal("Hello World from Go !!")
	})

}
