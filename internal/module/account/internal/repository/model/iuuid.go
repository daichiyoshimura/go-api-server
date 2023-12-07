package model

type iUUIDHelper interface {
	toBinary(id string) ([]byte, error)
	toString(id []byte) (string, error)
	generate() ([]byte, error)
}
