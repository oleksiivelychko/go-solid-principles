package main

type iBase interface {
	actionB()
}

type iDerivedBase interface {
	iBase
	actionA()
}

type derivedA struct{}

type derivedB struct{}

func (d *derivedA) actionA() {}

func (d *derivedA) actionB() {}

func (d *derivedB) actionB() {}
