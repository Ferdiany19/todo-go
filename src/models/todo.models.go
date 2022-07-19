package models

type Todo struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"title"`
	Description string `json:"description"`
}
