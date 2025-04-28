package jobcontroller

import (
	"job-search-manager/internal/dao"
	jobmodel "job-search-manager/internal/jobs/models"

	"gorm.io/gorm/clause"
)

type CompanyController struct{}

func (companyController *CompanyController) CreateACompany(company *jobmodel.Company) (err error) {

	result := dao.DB.Create(&company)

	return result.Error
}

func (companyController *CompanyController) GetAllCompanies() (companyList []*jobmodel.Company, err error) {
	result := dao.DB.Preload(clause.Associations).Find(&companyList)

	if result.Error != nil {
		return nil, result.Error
	}

	return companyList, result.Error
}

func (companyController *CompanyController) GetACompany(id string) (company *jobmodel.Company, err error) {
	company = new(jobmodel.Company)

	result := dao.DB.Preload(clause.Associations).Where(SQL_WHERE_ID, id).First(company)

	if result.Error != nil {
		return nil, result.Error
	}

	return company, result.Error
}

func (companyController *CompanyController) UpdateACompany(company *jobmodel.Company) (err error) {
	result := dao.DB.Begin().Save(company)

	return result.Error
}

func (companyController *CompanyController) DeleteACompany(id string) (err error) {
	result := dao.DB.Where(SQL_WHERE_ID, id).Delete(&jobmodel.Company{})

	return result.Error
}

func (companyController *CompanyController) AddJobtoCompany(company *jobmodel.Company, jobID string) (err error) {
	dao.DB.Model(&company).Association("Jobs").Append(&jobmodel.Job{ID: jobID})
	result := dao.DB.Model(&company).Update("Jobs", company.Jobs)

	return result.Error
}
