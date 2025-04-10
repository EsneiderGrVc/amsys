package api_services

import (
	"amsys/configs"
	orm_models "amsys/internal/api/models"
)

func GetVersionsByRepo(repo string) ([]orm_models.ControlVersion, bool) {
	var results []orm_models.ControlVersion
	if err := configs.DB.Where("repo_name = ?", repo).Find(&results).Error; err != nil {
		return nil, true
	}
	return results, false
}

func CreateVersion(version orm_models.ControlVersion) error {
	if err := configs.DB.Create(&version).Error; err != nil {
		return err
	}
	return nil
}
