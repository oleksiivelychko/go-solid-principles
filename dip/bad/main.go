package main

type serviceA struct{}

type serviceB struct {
	service *serviceA
}

func (e *serviceA) action() {}

// serviceB depends on serviceA (implementation)
func (s *serviceB) action() {
	s.service.action()
}
