package main

import "testing"

func TestDependencyInversion(t *testing.T) {
	s := &serviceC{&serviceA{}}
	if s.action() != "A" {
		t.Error("incorrect implementation")
	}

	s = &serviceC{&serviceB{}}
	if s.action() != "B" {
		t.Error("incorrect implementation")
	}
}
