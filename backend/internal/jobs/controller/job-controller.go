package jobcontroller

import (
	"job-search-manager/internal/dao"
	jobmodel "job-search-manager/internal/jobs/models"
)

func CreateAJob(job *jobmodel.Job) (err error) {

	job.DefaultValues()

	result := dao.DB.Create(&job)

	return result.Error
}

func GetAllJob() (jobList []*jobmodel.Job, err error) {
	result := dao.DB.Find(&jobList)

	if result.Error != nil {
		return nil, result.Error
	}

	return jobList, result.Error
}

func GetAJob(id string) (job *jobmodel.Job, err error) {
	job = new(jobmodel.Job)

	result := dao.DB.Where(SQL_WHERE_ID, id).First(job)

	if result.Error != nil {
		return nil, result.Error
	}

	return job, result.Error
}

func UpdateAJob(job *jobmodel.Job) (err error) {
	result := dao.DB.Save(job)

	return result.Error
}

func DeleteAJob(id string) (err error) {
	result := dao.DB.Where(SQL_WHERE_ID, id).Delete(&jobmodel.Job{})

	return result.Error
}
