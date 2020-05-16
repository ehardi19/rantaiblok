package repository

import (
	"github.com/ehardi19/rantaiblok/model"
	"github.com/jinzhu/gorm"
)

// SaveBlock inserts Block to database (Node)
func (repo *databaseRepository) SaveBlock(block model.Block) (err error) {
	err = repo.DB.Save(&block).Error

	return
}

// GetAllBlock gets all Block from database
func (repo *databaseRepository) GetAllBlock() (blocks []model.Block, err error) {
	err = repo.DB.Find(&blocks).Error

	return
}

// GetLastBlock gets last Block from database
func (repo *databaseRepository) GetLastBlock() (block model.Block, err error) {
	err = repo.DB.Last(&block).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return model.Block{}, nil
	}

	return
}

// GetBlockByID gets Block by ID (int) from database
func (repo *databaseRepository) GetBlockByID(id int) (block model.Block, err error) {
	err = repo.DB.Where("id = ?", id).Find(&block).Error

	return
}

// Count gets length info of Blockchain from database
func (repo *databaseRepository) Count() (count int, err error) {
	err = repo.DB.Model(&model.Block{}).Count(&count).Error

	return
}
