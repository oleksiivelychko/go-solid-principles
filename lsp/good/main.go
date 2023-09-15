package main

type iBase interface {
	actionB() string
}

type iDerivedBase interface {
	iBase
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
