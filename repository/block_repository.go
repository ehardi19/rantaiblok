package repository

import (
	"github.com/ehardi19/rantaiblok/model"
	"github.com/jinzhu/gorm"
)

func (repo *databaseRepository) SaveBlock(block model.Block) (err error) {
	err = repo.DB.Save(&block).Error

	return
}

func (repo *databaseRepository) GetAllBlock() (blocks []model.Block, err error) {
	err = repo.DB.Find(&blocks).Error

	return
}

func (repo *databaseRepository) GetLastBlock() (block model.Block, err error) {
	err = repo.DB.Last(&block).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return model.Block{}, nil
	}

	return
}

func (repo *databaseRepository) GetBlockByID(id int) (block model.Block, err error) {
	err = repo.DB.Where("id = ?", id).Find(&block).Error

	return
}

func (repo *databaseRepository) Count() (count int, err error) {
	err = repo.DB.Model(&model.Block{}).Count(&count).Error

	return
}
