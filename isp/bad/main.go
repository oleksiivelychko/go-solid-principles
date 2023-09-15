package main

type iEntity interface {
	actionA()
	actionB()
	actionC()
}

type entityA struct{}

type entityB struct{}

func (e *entityA) actionA() {}

func (e *entityA) actionB() {}

func (e *entityA) actionC() {}

func (e *entityB) actionA() {}

func (e *entityB) actionB() {
	// not supported
}

func (e *entityB) actionC() {
	// not supported
}
