package main

// entity has only one responsibility: manage data
type entity struct{}

func (e *entity) data() {}

// entityRepository has only one responsibility: save data
type entityRepository struct{}

func (r *entityRepository) save(e *entity) {}
