package main

type service interface{ action() string }

type serviceA struct{}
type serviceB struct{}

// serviceC depends on service (abstraction)
type serviceC struct {
	service service
}

func (s *serviceA) action() string { return "A" }
func (s *serviceB) action() string { return "B" }

func (s *serviceC) action() string {
	return s.service.action()
}
