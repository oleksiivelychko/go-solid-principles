package srp

/*
Single responsibility.
A class should have one, and only one, reason to change. (Robert C Martin)
*/

/*
entity has only one responsibility to manage data.
*/
type entity struct {
	name string
}

func (e *entity) getName() string {
	return e.name
}

/*
entityRepository has only one responsibility to handle DB operations.
*/
type entityRepository struct {
	// database connection or other storage-related fields
}

func (repo *entityRepository) save(e *entity) error {
	// save entity to the database
	return nil
}
