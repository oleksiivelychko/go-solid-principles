package main

import "testing"

func Test(t *testing.T) {
	service := &serviceC{service: &serviceA{}}
	if service.actionB() != "A" {
		t.Error("got incorrect service implementation")
	}

	service = &serviceC{service: &serviceB{}}
	if service.actionB() != "B" {
		t.Error("got incorrect service implementation")
	}
}
