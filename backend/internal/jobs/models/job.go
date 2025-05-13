package jobmodel

import (
	"strconv"
	"time"
)

type Job struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Date        string    `json:"date"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Notes       string    `json:"notes"`
	Status      JobStatus `json:"status"`
	CompanyID   string
}

func (job *Job) DefaultValues() {
	timestamp := time.Now().Unix()

	job.Status = PENDING
	job.Date = strconv.FormatInt(timestamp, 10)
}

func MakeJob(id string, title string, desc string, notes string) *Job {
	job := Job{
		ID:          id,
		Title:       title,
		Description: desc,
		Notes:       notes,
	}
	job.DefaultValues()

	return &job
}
