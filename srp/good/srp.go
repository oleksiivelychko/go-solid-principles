package good

/*
Single responsibility in right way.
*/

/*
entity has only one responsibility: managing data.
*/
type entity struct{}

func (e *entity) getData() {
	// managing data
}

/*
entityRepository has only one responsibility: saving data.
*/
type entityRepository struct {
	// database connection or other storage-related fields
}

func (repo *entityRepository) save(e *entity) {
	// saving data
}
