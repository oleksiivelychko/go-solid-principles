package main

type base interface {
	actionA() string
}

type derivedA struct{}

func (d *derivedA) actionA() string {
	return "derivedA is doing actionA."
}

// derivedB is not intended for actionA() and violates the principle.
type derivedB struct{}

func (d *derivedB) actionA() string {
	return "derivedB is doing actionA."
}
