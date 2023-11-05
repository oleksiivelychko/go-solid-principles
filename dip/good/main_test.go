package main

import "testing"

func TestDependencyInversion(t *testing.T) {
	var (
		a = &serviceC{&serviceA{}}
		b = &serviceC{&serviceB{}}
	)

	g := a.action()
	if g != "A" {
		t.Errorf("expected %q, got %q", "A", g)
	}

	g = b.action()
	if g != "B" {
		t.Errorf("expected %q, got %q", "B", g)
	}
}
