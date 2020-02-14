package blockchain

import (
	"github.com/ehardi19/rantaiblok/models"
)

// Repository ..
type Repository interface {
	Fetch() ([]models.Block, error)
	GetByID(id int64) (models.Block, error)
	Store(block models.Block) error
}
