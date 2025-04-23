package jobmodel

import (
	"strconv"
	"time"
)

type JobStatus string

const (
	Pending   JobStatus = "pending"
	Rejected  JobStatus = "rejected"
	Contacted JobStatus = "contacted"
	Offer     JobStatus = "offer"
)

type Job struct {
	ID          string    `json:"id"`
	Date        string    `json:"date"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Notes       string    `json:"notes"`
	Status      JobStatus `json:"status"`
}

func (job *Job) DefaultValues() {
	timestamp := time.Now().Unix()

	job.Status = Pending
	job.Date = strconv.FormatInt(timestamp, 10)
}
