package main

// entity has only one responsibility: manage data
type entity struct{}

func (e *entity) data() {}

// repository has only one responsibility: save data
type repository struct{}

func (r *repository) save(e *entity) {}
