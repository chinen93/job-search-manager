package jobmodel

type Company struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Jobs []Job  `json:"jobs" gorm:"foreignKey:CompanyID"`
}
