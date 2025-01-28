package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password []byte `json:"-"` // dont return password
}