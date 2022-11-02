package models

type User struct {
	Id       int    `gorm:"primary key;autoincrement" json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Age      string `json:"age"`
}
