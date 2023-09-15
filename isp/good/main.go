package main

type iEntityA interface {
	actionA()
}

type iEntityB interface {
	actionB()
}

type iEntityC interface {
	actionC()
}

type entityA struct{}

type entityB struct{}

func (e *entityA) actionA() {}

func (e *entityA) actionB() {}

func (e *entityA) actionC() {}

// implement only necessary methods
func (e *entityB) actionA() {}
