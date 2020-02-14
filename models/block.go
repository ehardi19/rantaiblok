package models

// BlockRequest ..
type BlockRequest struct {
	Data string `json:"data"`
}

// Block ..
type Block struct {
	ID        int64  `json:"id" gorm:"primary_key" gorm:"unique"`
	Data      string `json:"data"`
	Timestamp string `json:"timestamp"`
	Hash      string `json:"hash" gorm:"unique"`
	PrevHash  string `json:"prev_hash" gorm:"unique"`
}
