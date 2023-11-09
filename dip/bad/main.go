package main

type serviceA struct{}

type serviceB struct {
	service *serviceA
}

func (s *serviceA) action() {}

// serviceB depends on serviceA implementation
func (s *serviceB) action() { s.service.action() }
