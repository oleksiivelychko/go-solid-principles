package main

import "testing"

func TestDependencyInversion(t *testing.T) {
	var (
		a = &serviceC{&serviceA{}}
		b = &serviceC{&serviceB{}}
	)

	res := a.action()
	if res != "A" {
		t.Errorf("expected %q, got %q", "A", res)
	}

	res = b.action()
	if res != "B" {
		t.Errorf("expected %q, got %q", "B", res)
	}
}
