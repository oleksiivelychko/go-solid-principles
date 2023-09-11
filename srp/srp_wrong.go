package srp

/*
Single responsibility in wrong way.
*/

// wrongEntity has two responsibilities: data management and data storage.
type wrongEntity struct{}

func (e *wrongEntity) getData() {
	// data management
}

func (e *wrongEntity) save() {
	// data storage
}
