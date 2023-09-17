package main

// entity has two responsibilities: managing data and saving data.
type entity struct{}

// managing data
func (e *entity) getData() {}

// saving data
func (e *entity) save() {}
