package env

type DB struct {
	host     string
	user     string
	password string
	instance string
}

func (e *DB) Host() string {
	return e.host
}

func (e *DB) User() string {
	return e.user
}

func (e *DB) Password() string {
	return e.password
}

func (e *DB) Instance() string {
	return e.instance
}
