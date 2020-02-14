package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/ehardi19/rantaiblok/blockchain"
	"github.com/ehardi19/rantaiblok/models"
)

// Usecase ..
type Usecase struct {
	Repository blockchain.Repository
}

// NewUsecase ..
func NewUsecase(repo blockchain.Repository) blockchain.Usecase {
	return &Usecase{
		Repository: repo,
	}
}

// Fetch ..
func (u *Usecase) Fetch() ([]models.Block, error) {
	blockchain, err := u.Repository.Fetch()
	if err != nil {
		return []models.Block{}, err
	}

	return blockchain, nil
}

// GetByID ..
func (u *Usecase) GetByID(id int64) (models.Block, error) {
	block, err := u.Repository.GetByID(id)
	if err != nil {
		return models.Block{}, err
	}

	return block, nil
}

// Store ..
func (u *Usecase) Store(req models.BlockRequest) (models.Block, error) {
	blockchain, err := u.Fetch()
	if err != nil {
		return models.Block{}, err
	}

	prevID := int64(len(blockchain))
	prevHash := blockchain[prevID-1].Hash

	block := models.Block{
		ID:        prevID + 1,
		Data:      req.Data,
		Timestamp: time.Now().String(),
		PrevHash:  prevHash,
	}

	hash := Hash(block)
	block.Hash = hash

	err = u.Repository.Store(block)
	if err != nil {
		return models.Block{}, err
	}

	return block, nil
}

// Hash ..
func Hash(block models.Block) string {
	var blockString string
	blockString = strconv.FormatInt(block.ID, 10) + block.Data + block.Timestamp + block.PrevHash
	hashed := sha256.Sum256([]byte(blockString))

	return hex.EncodeToString(hashed[:])
}

// Validate ..
func (u *Usecase) Validate() bool {
	blockchain, _ := u.Fetch()

	for i := 1; i < len(blockchain)-1; i++ {
		block := blockchain[i]
		nextBlock := blockchain[i+1]

		hashedBlock := Hash(block)

		if block.Hash != hashedBlock {
			return false
		}

		if block.Hash != nextBlock.PrevHash {
			return false
		}
	}

	return true
}
