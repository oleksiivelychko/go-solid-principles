package main

type entityA struct {
	paramA int
	paramB int
}

type entityB struct {
	paramC int
}

// open for expansion (new types of iEntity might be implemented through interface)
// closed for modification (doesn't need to change the total() function after adding each new iEntity).
type iEntity interface {
	total() int
}

func (e *entityA) total() int {
	return e.paramA * e.paramB
}

func (e *entityB) total() int {
	return e.paramC * 2
}
