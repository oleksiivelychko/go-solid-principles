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

func (a *serviceA) actionA() string {
	return "A"
}

func (b *serviceB) actionA() string {
	return "B"
}

func (c *serviceC) actionB() string {
	return c.service.actionA()
}
