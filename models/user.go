package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
