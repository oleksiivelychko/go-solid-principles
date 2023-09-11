package srp

/*
Single responsibility [GOOD].
*/

/*
goodEntity has only one responsibility: data management.
*/
type goodEntity struct{}

func (e *goodEntity) getData() {
	// data management
}

/*
entityRepository has only one responsibility: data storage.
*/
type entityRepository struct {
	// database connection
}

func (repo *entityRepository) save(e *goodEntity) {
	// data storage
}
