package main

type actionA interface{ actionA() }
type actionB interface{ actionB() }
type actionC interface{ actionC() }

type entityA struct{}
type entityB struct{}

func (e *entityA) actionA() {}
func (e *entityA) actionB() {}
func (e *entityA) actionC() {}

// implement only necessary methods
func (e *entityB) actionA() {}
