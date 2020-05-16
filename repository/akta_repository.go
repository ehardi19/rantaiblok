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

func (repo *databaseRepository) GetAktaByID(id int) (akta model.Akta, err error) {
	err = repo.DB.Where("id = ?", id).Find(&akta).Error

	return
}

func (repo *databaseRepository) GetAktaByAktaNum(aktaNum string) (akta model.Akta, err error) {
	err = repo.DB.Where("akta_number = ?", aktaNum).Find(&akta).Error

	return
}
