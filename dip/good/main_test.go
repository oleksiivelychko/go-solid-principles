package main

import "testing"

func TestDependencyInversion(t *testing.T) {
	s := &serviceC{&serviceA{}}
	got := s.action()

	if got != "A" {
		t.Errorf("expected %q, got %q", "A", got)
	}

	s = &serviceC{&serviceB{}}
	got = s.action()

	if got != "B" {
		t.Errorf("expected %q, got %q", "B", got)
	}
}
