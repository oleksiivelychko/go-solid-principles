package main

type classA interface {
	classB
	actionA()
}

type classB interface{ actionB() }

type subclassA struct{}
type subclassB struct{}

func (s *subclassA) actionA() {}
func (s *subclassA) actionB() {}
func (s *subclassB) actionB() {}
