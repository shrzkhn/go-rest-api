package models

import "github.com/google/uuid"

type City struct {
	Id      uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name    string    `json:"name"`
	Country string    `json:"country"`
}
