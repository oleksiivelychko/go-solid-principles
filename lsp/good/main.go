package main

type base interface {
	actionB() string
}

type derivedBase interface {
	base
	actionA() string
}

type derivedA struct{}

func (d *derivedA) actionA() string {
	return "derivedA is doing actionA."
}

func (d *derivedA) actionB() string {
	return "derivedB is doing actionB."
}

type derivedB struct{}

func (d *derivedB) actionB() string {
	return "derivedB is doing actionB."
}
