package srp

/*
Single responsibility [BAD].
*/

// badEntity has two responsibilities: data management and data storage.
type badEntity struct{}

func (e *badEntity) getData() {
	// data management
}

func (e *badEntity) save() {
	// data storage
}
