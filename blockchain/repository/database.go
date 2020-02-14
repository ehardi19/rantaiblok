package repository

import (
	"github.com/ehardi19/rantaiblok/blockchain"
	"github.com/ehardi19/rantaiblok/models"
	"github.com/jinzhu/gorm"
)

type gormRepository struct {
	DB *gorm.DB
}

// NewGormRepository ..
func NewGormRepository(DB *gorm.DB) blockchain.Repository {
	DB.AutoMigrate(
		&models.Block{},
	)

	return &gormRepository{DB: DB}
}

// Fetch ..
func (g *gormRepository) Fetch() ([]models.Block, error) {
	var blockchain []models.Block
	if err := g.DB.Model(&blockchain).Find(&blockchain).Error; err != nil {
		return nil, err
	}

	return blockchain, nil
}

// GetByID ..
func (g *gormRepository) GetByID(id int64) (models.Block, error) {
	var block models.Block
	if err := g.DB.Model(&block).Where("id = ?", id).Find(&block).Error; err != nil {
		return models.Block{}, err
	}

	return block, nil
}

// Store ..
func (g *gormRepository) Store(block models.Block) error {
	if err := g.DB.Model(&block).Create(&block).Error; err != nil {
		return err
	}

	return nil
}
