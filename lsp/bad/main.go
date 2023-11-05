package main

type class interface{ actionA() }

type classA struct{}
type classB struct{}

func (s *classA) actionA() {}

// classB is not intended for actionA() and violates the principle
func (s *classB) actionA() {}
