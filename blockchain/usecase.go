package blockchain

import (
	"github.com/ehardi19/rantaiblok/models"
)

// Usecase ..
type Usecase interface {
	Fetch() ([]models.Block, error)
	GetByID(id int64) (models.Block, error)
	Store(req models.BlockRequest) (models.Block, error)
	Validate() bool
}
