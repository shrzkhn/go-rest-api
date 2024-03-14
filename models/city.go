package models

type City struct {
	Id      string `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Country string `json:"country"`
}
