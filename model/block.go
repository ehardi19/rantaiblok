package model

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

// BlockData ...
type BlockData interface {
	ToJSON() (string, error)
}

type CreateBlockRequest struct {
	Data string `json:"data"`
}

// Block ..
type Block struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Timestamp string `json:"timestamp"`
	Nonce     int    `json:"nonce"`
	PrevHash  string `json:"prev_hash"`
	Data      string `json:"data"`
	Hash      string `json:"hash"`
}

// GenerateHash ...
func (block *Block) GenerateHash() string {
	data := make(map[string]interface{})
	data["id"] = block.ID
	data["timestamp"] = block.Timestamp
	data["nonce"] = block.Nonce
	data["prev_hash"] = block.PrevHash
	data["data"] = block.Data

	raw, _ := json.Marshal(data)

	h := sha256.New()
	h.Write(raw)
	hash := hex.EncodeToString(h.Sum(raw))

	return hash
}

// GenerateNewBlock ...
func GenerateNewBlock(id int, timestamp, prevHash, data string) (*Block, error) {

	var nonce int = 0
	newBlock := Block{
		id,
		timestamp,
		nonce,
		prevHash,
		data,
		"",
	}

	var hash string
	for hash = newBlock.GenerateHash(); len(hash) >= 3 && hash[len(hash)-3:] != "ace"; hash = newBlock.GenerateHash() {
		nonce++
		newBlock = Block{
			id,
			timestamp,
			nonce,
			prevHash,
			data,
			hash,
		}
	}

	return &Block{
		id,
		timestamp,
		nonce,
		prevHash,
		data,
		hash,
	}, nil
}
