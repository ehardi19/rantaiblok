package service

import (
	"errors"
	"strings"
	"time"

	"github.com/ehardi19/rantaiblok/model"
	"github.com/sirupsen/logrus"
)

// IsValid checks validity of blockchain
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

// SaveBlock makes new block and saves it to all nodes
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

// GetAllBlock gets all block from blockchain
func (s *Service) GetAllBlock() ([]model.Block, error) {
	blocks, err := s.Node1.GetAllBlock()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return blocks, err
}

// GetLastBlock gets last block of blockchain
func (s *Service) GetLastBlock() (model.Block, error) {
	block, err := s.Node1.GetLastBlock()
	if err != nil {
		logrus.Error(err)
		return model.Block{}, err
	}

	return block, err
}

// GetBlockByID gets block by ID (int) of blockchain
func (s *Service) GetBlockByID(id int) (model.Block, error) {
	block, err := s.Node1.GetBlockByID(id)
	if err != nil {
		logrus.Error(err)
		return model.Block{}, err
	}

	return block, err
}

// PushDataToBlock process data in data pool to blockchain nodes
// In this case every 3 data in pool created into 1 block data
func (s *Service) PushDataToBlock() error {
	// Fetch Pool Database
	aktas, err := s.Pool.GetAllAkta()
	if err != nil {
		logrus.Error(err)
		return err
	}

	// Rule: 3 Data Per Block
	if len(aktas) < 3 {
		return errors.New("too few data to process")
	}

	arrData := []string{}
	for i := 0; i < 3; i++ {
		jsonData, err := aktas[i].ToJSON()
		if err != nil {
			return err
		}

		arrData = append(arrData, jsonData)
	}
	strData := strings.Join(arrData, ",")

	// Create Block
	err = s.SaveBlock(model.CreateBlockRequest{Data: strData})
	if err != nil {
		logrus.Error(err)
		return err
	}

	// Delete Pushed Data From Pool
	for i := 0; i < 3; i++ {
		err = s.Pool.DeleteAktaByID(aktas[i].ID)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	return nil
}
