package bad

/*
Single responsibility in wrong way.
*/

// entity has two responsibilities: managing data and saving data.
type entity struct{}

func (e *entity) getData() {
	// managing data
}

func (e *entity) save() {
	// saving data
}
