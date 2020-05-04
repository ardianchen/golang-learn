package models

type crud interface {
	create() (bool, error)
	read() (bool, error)
	update() (bool, error)
	delete() (bool, error)
}
