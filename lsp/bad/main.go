package main

type iBase interface {
	actionA()
}

type derivedA struct{}

type derivedB struct{}

func (d *derivedA) actionA() {}

// derivedB is not intended for actionA() and violates the principle.
func (d *derivedB) actionA() {}
