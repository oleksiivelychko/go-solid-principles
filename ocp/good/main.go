package main

// open for expansion (new types of entity might be implemented through interface)
// closed for modification (doesn't need to change the total() after adding each new entity)
type entity interface{ total() int }

type entityA struct {
	a int
	b int
}

type entityB struct {
	c int
}

func (e *entityA) total() int { return e.a * e.b }
func (e *entityB) total() int { return e.c * 2 }
