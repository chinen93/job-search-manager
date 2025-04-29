package jobcontroller

import (
	"job-search-manager/internal/dao"
	jobmodel "job-search-manager/internal/jobs/models"

	"gorm.io/gorm/clause"
)

type CompanyController struct{}

func (companyController *CompanyController) Create(company *jobmodel.Company) (err error) {

	result := dao.DB.Create(&company)

	return result.Error
}

func (companyController *CompanyController) GetAll() (companyList []*jobmodel.Company, err error) {
	result := dao.DB.Preload(clause.Associations).Find(&companyList)

	if result.Error != nil {
		return nil, result.Error
	}

	return companyList, result.Error
}

func (companyController *CompanyController) Get(id string) (company *jobmodel.Company, err error) {
	company = new(jobmodel.Company)

	result := dao.DB.Preload(clause.Associations).Where(SQL_WHERE_ID, id).First(company)

	if result.Error != nil {
		return nil, result.Error
	}

	return company, result.Error
}

func (companyController *CompanyController) Update(company *jobmodel.Company) (err error) {
	result := dao.DB.Begin().Save(company)

	return result.Error
}

func (companyController *CompanyController) Delete(id string) (err error) {
	result := dao.DB.Where(SQL_WHERE_ID, id).Delete(&jobmodel.Company{})

	return result.Error
}

func (companyController *CompanyController) AddJob(company *jobmodel.Company, job *jobmodel.Job) (err error) {
	dao.DB.Model(&company).Association("Jobs").Append(job)
	result := dao.DB.Model(&company).Update("Jobs", company.Jobs)

	return result.Error
}
