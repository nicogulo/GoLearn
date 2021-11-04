package models

type User struct {
	Id       uint   `json:"id" validate:"require"`
	Name     string `json:"name" validate:"require"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}
