package service

import (
	"time"

	"github.com/ehardi19/rantaiblok/model"
	"github.com/sirupsen/logrus"
)

// IsValid ...
func (s *Service) IsValid() (bool, error) {
	// Checking validity of Node 1
	count1, err := s.Node1.Count()
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	node1, err := s.Node1.GetAllBlock()
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	for i := 1; i < count1; i++ {
		prevBlock := node1[i-1]
		block := node1[i]

		hashedBlock := block.GenerateHash()

		if block.Hash != hashedBlock {
			return false, err
		}

		if block.PrevHash != prevBlock.Hash {
			return false, err
		}
	}

	// Checking validity of Node 2
	node2, err := s.Node2.GetAllBlock()
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	count2, err := s.Node2.Count()
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	for i := 1; i < count2; i++ {
		prevBlock := node2[i-1]
		block := node2[i]

		hashedBlock := block.GenerateHash()

		if block.Hash != hashedBlock {
			return false, err
		}

		if block.PrevHash != prevBlock.Hash {
			return false, err
		}
	}

	// Checking validity of Node 3
	node3, err := s.Node3.GetAllBlock()
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	count3, err := s.Node3.Count()
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	for i := 1; i < count3; i++ {
		prevBlock := node3[i-1]
		block := node3[i]

		hashedBlock := block.GenerateHash()

		if block.Hash != hashedBlock {
			return false, err
		}

		if block.PrevHash != prevBlock.Hash {
			return false, err
		}
	}

	// Checking length
	if !(count1 == count2 && count2 == count3) {
		return false, nil
	}

	// Checking data between nodes
	for i := 0; i < count1; i++ {
		if !(node1[i] == node2[i] && node2[i] == node3[i]) {
			return false, nil
		}
	}

	return true, nil

}

// SaveBlock ...
func (s *Service) SaveBlock(req model.CreateBlockRequest) error {
	count, err := s.Node1.Count()
	if err != nil {
		return err
	}

	lastBlock, err := s.Node1.GetLastBlock()
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	prevID := count
	prevHash := lastBlock.Hash
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

	// Saving to Node1
	err = s.Node1.SaveBlock(*block)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// Saving to Node2
	err = s.Node2.SaveBlock(*block)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// Saving to Node3
	err = s.Node3.SaveBlock(*block)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

// GetAllBlock ...
func (s *Service) GetAllBlock() ([]model.Block, error) {
	blocks, err := s.Node1.GetAllBlock()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return blocks, err
}

// GetLastBlock ...
func (s *Service) GetLastBlock() (model.Block, error) {
	block, err := s.Node1.GetLastBlock()
	if err != nil {
		logrus.Error(err)
		return model.Block{}, err
	}

	return block, err
}

// GetBlockByID ...
func (s *Service) GetBlockByID(id int) (model.Block, error) {
	block, err := s.Node1.GetBlockByID(id)
	if err != nil {
		logrus.Error(err)
		return model.Block{}, err
	}

	return block, err
}
