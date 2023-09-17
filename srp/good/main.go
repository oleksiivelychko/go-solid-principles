package main

// entity has only one responsibility: managing data.
type entity struct{}

// entityRepository has only one responsibility: saving data.
type entityRepository struct{}

// managing data
func (e *entity) getData() {}

// saving data
func (repo *entityRepository) save(e *entity) {}
