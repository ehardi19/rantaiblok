package repository

import (
	"fmt"
	"os"

	"github.com/ehardi19/rantaiblok/model"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Repository ..
type Repository interface {
	SaveBlock(block model.Block) error
	GetAllBlock() ([]model.Block, error)
	GetLastBlock() (model.Block, error)
	GetBlockByID(id int) (model.Block, error)
	Count() (int, error)

	SaveAkta(akta model.Akta) error
	GetAllAkta() ([]model.Akta, error)
	GetAktaByID(id int) (model.Akta, error)
	GetAktaByAktaNum(aktaNum string) (model.Akta, error)
}

type databaseRepository struct {
	*gorm.DB
}

// Init ..
func Init() Repository {
	godotenv.Load()
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASSWORD")

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	db.AutoMigrate(&model.Block{})
	db.AutoMigrate(&model.Akta{})

	return &databaseRepository{db}
}
