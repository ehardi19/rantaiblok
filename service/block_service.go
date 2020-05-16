package service

import (
	"time"

	"github.com/ehardi19/rantaiblok/model"
	"github.com/sirupsen/logrus"
)

// IsValid ...
func (s *Service) IsValid() (bool, error) {
	blocks, err := s.Repo.GetAllBlock()
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	for i := 1; i < len(blocks)-1; i++ {
		block := blocks[i]
		nextBlock := blocks[i+1]

		hashedBlock := block.GenerateHash()

		if block.Hash != hashedBlock {
			return false, err
		}

		if block.Hash != nextBlock.PrevHash {
			return false, err
		}
	}

	return true, nil
}

// SaveBlock ...
func (s *Service) SaveBlock(req model.CreateBlockRequest) error {
	count, err := s.Repo.Count()
	if err != nil {
		return err
	}

	lastBlock, err := s.Repo.GetLastBlock()
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	prevID := count
	prevHash := lastBlock.PrevHash
	timestamp := time.Now().String()

	block, err := model.GenerateNewBlock(
		prevID+1,
		timestamp,
		prevHash,
		req.Data,
	)

	if err != nil {
		logrus.Error(err)
		return err
	}

	err = s.Repo.SaveBlock(*block)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

// GetAllBlock ...
func (s *Service) GetAllBlock() ([]model.Block, error) {
	blocks, err := s.Repo.GetAllBlock()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	logrus.Info(blocks)

	return blocks, err
}

// GetLastBlock ...
func (s *Service) GetLastBlock() (model.Block, error) {
	block, err := s.Repo.GetLastBlock()
	if err != nil {
		logrus.Error(err)
		return model.Block{}, err
	}

	logrus.Info(block)

	return block, err
}

// GetBlockByID ...
func (s *Service) GetBlockByID(id int) (model.Block, error) {
	block, err := s.Repo.GetBlockByID(id)
	if err != nil {
		logrus.Error(err)
		return model.Block{}, err
	}

	logrus.Info(block)

	return block, err
}
