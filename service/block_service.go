package service

import (
	"time"

	"github.com/ehardi19/rantaiblok/model"
	"github.com/sirupsen/logrus"
)

// IsValid ...
func (s *Service) IsValid() (bool, error) {
	// Checking validity of Node 1
	node1, err := s.Node1.GetAllBlock()
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	for i := 1; i < len(node1)-1; i++ {
		block := node1[i]
		nextBlock := node1[i+1]

		hashedBlock := block.GenerateHash()

		if block.Hash != hashedBlock {
			return false, err
		}

		if block.Hash != nextBlock.PrevHash {
			return false, err
		}
	}

	// Checking validity of Node 2
	node2, err := s.Node2.GetAllBlock()
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	for i := 1; i < len(node2)-1; i++ {
		block := node2[i]
		nextBlock := node2[i+1]

		hashedBlock := block.GenerateHash()

		if block.Hash != hashedBlock {
			return false, err
		}

		if block.Hash != nextBlock.PrevHash {
			return false, err
		}
	}

	// Checking validity of Node 3
	node3, err := s.Node3.GetAllBlock()
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	for i := 1; i < len(node3)-1; i++ {
		block := node3[i]
		nextBlock := node3[i+1]

		hashedBlock := block.GenerateHash()

		if block.Hash != hashedBlock {
			return false, err
		}

		if block.Hash != nextBlock.PrevHash {
			return false, err
		}
	}

	if !(len(node1) == len(node2) && len(node2) == len(node3)) {
		return false, nil
	}

	for i := 0; i < len(node1); i++ {
		if !(node1[i] == node2[i] && node2[i] == node3[i]) {
			logrus.Println(node1[i])
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
