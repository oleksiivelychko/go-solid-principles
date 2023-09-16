package main

type iService interface {
	actionA() string
}

type serviceA struct{}

type serviceB struct{}

// serviceC depends on iService (abstraction)
type serviceC struct {
	service iService
}

func (s *serviceA) actionA() string {
	return "A"
}

func (s *serviceB) actionA() string {
	return "B"
}

func (s *serviceC) actionB() string {
	return s.service.actionA()
}
