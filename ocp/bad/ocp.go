package ocp_bad

type entityA struct {
	paramA int
	paramB int
}

type entityB struct {
	paramC int
}

// need to modify implementation for each new entity.
func total(object interface{}) int {
	switch s := object.(type) {
	case *entityA:
		return s.paramA * s.paramB
	case *entityB:
		return s.paramC * 2
	default:
		return 0
	}
}
