package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	name     string `gorm:"type:varchar(50)" validate:"required"`
	password string `gorm:"type:varchar(50)" `
	email    string `gorm:"type:varchar(50)" validate:"required,email"`
	// Avatar   string
}

func (this User) update() (bool, error) {
	return false, nil
}
func (this User) read() (bool, error) {
	return false, nil
}
func (this User) create() (bool, error) {
	return false, nil
}
func (this User) delete() (bool, error) {
	return false, nil
}
