package repository

import (
	"github.com/ehardi19/rantaiblok/model"
)

// SaveAkta inserts Akta to database (in this case is Data Pool)
func (repo *databaseRepository) SaveAkta(akta model.Akta) (err error) {
	err = repo.Save(&akta).Error

	return
}

// GetAllAkta gets all Akta from database
func (repo *databaseRepository) GetAllAkta() (aktas []model.Akta, err error) {
	err = repo.DB.Find(&aktas).Error

	return
}

// GetAktaByID gets an Akta by ID (int) from database
func (repo *databaseRepository) GetAktaByID(id int) (akta model.Akta, err error) {
	err = repo.DB.Where("id = ?", id).Find(&akta).Error

	return
}

// DeleteAktaByID deletes an Akta by ID (int) from database
func (repo *databaseRepository) DeleteAktaByID(id int) (err error) {
	var akta model.Akta
	err = repo.DB.Where("id = ?", id).Delete(&akta).Error

	return
}
