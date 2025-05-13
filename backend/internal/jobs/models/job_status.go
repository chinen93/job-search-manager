package jobmodel

type JobStatus string

const (
	PENDING   JobStatus = "pending"
	REJECTED  JobStatus = "rejected"
	CONTACTED JobStatus = "contacted"
	OFFER     JobStatus = "offer"
)
