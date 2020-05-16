package repository

import (
	"github.com/ehardi19/rantaiblok/model"
)

func (repo *databaseRepository) SaveAkta(akta model.Akta) (err error) {
	err = repo.Save(&akta).Error

	return
}

func (repo *databaseRepository) GetAllAkta() (aktas []model.Akta, err error) {
	err = repo.DB.Find(&aktas).Error

	return
}

func (repo *databaseRepository) QueryAllFakta(offset, limit int) (aktas []model.Akta, err error) {
	err = repo.DB.Offset(offset).Limit(limit).Find(&aktas).Error

	return
}

func (repo *databaseRepository) GetAktaByID(id int) (akta model.Akta, err error) {
	err = repo.DB.Where("id = ?", id).Find(&akta).Error

	return
}

func (repo *databaseRepository) DeleteAktaByID(id int) (err error) {
	var akta model.Akta
	err = repo.DB.Where("id = ?", id).Delete(&akta).Error

	return
}
