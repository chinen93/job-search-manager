package jobmodel

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	ID          string `json:"id"`
	Date        string `json:"date"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Notes       string `json:"notes"`
}
