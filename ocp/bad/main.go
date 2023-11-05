package main

type entityA struct {
	a int
	b int
}

type entityB struct{ c int }

// need to modify implementation for each new entity
func total(e interface{}) int {
	switch t := e.(type) {
	case *entityA:
		return t.a * t.b
	case *entityB:
		return t.c * 2
	default:
		return 0
	}
}
