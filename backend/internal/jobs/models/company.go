package jobmodel

import (
	"github.com/google/uuid"
)

type Company struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Jobs []Job  `json:"jobs" gorm:"foreignKey:CompanyID"`
}

func (company *Company) DefaultValues() {

	company.ID = uuid.NewString()
}
