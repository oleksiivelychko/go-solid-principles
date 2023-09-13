package ocp

/*
Open-closed in right way.
*/

type entityC struct {
	paramA int
	paramB int
}

type entityD struct {
	paramC int
}

/*
open for expansion (new types of entity might be implemented through interface)
closed for modification (doesn't need to change the totalCount() function after each new entity was added).
*/
type entity interface {
	totalCount() int
}

func (e *entityC) totalCount() int {
	return e.paramA * e.paramB
}

func (e *entityD) totalCount() int {
	return e.paramC * 2
}
