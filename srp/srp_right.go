package srp

/*
Single responsibility in right way.
*/

/*
rightEntity has only one responsibility: data management.
*/
type rightEntity struct{}

func (e *rightEntity) getData() {
	// data management
}

/*
entityRepository has only one responsibility: data storage.
*/
type entityRepository struct {
	// database connection
}

func (repo *entityRepository) save(e *rightEntity) {
	// data storage
}
