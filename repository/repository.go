package repository

import (
	"fmt"
	"os"

	"github.com/ehardi19/rantaiblok/model"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Repository defines repository and its available queries
type Repository interface {
	SaveBlock(block model.Block) error
	GetAllBlock() ([]model.Block, error)
	GetLastBlock() (model.Block, error)
	GetBlockByID(id int) (model.Block, error)
	Count() (int, error)

	SaveAkta(akta model.Akta) error
	GetAllAkta() ([]model.Akta, error)
	GetAktaByID(id int) (model.Akta, error)
	DeleteAktaByID(id int) error
}

// databaseRepository implements repository that has GORM (object relation mapping) of database
type databaseRepository struct {
	*gorm.DB
}

// InitNode1 represents database configuration of Node1
func InitNode1() Repository {
	godotenv.Load()
	dbHost := os.Getenv("NODE1_HOST")
	dbPort := os.Getenv("NODE1_PORT")
	dbUser := os.Getenv("NODE1_USER")
	dbName := os.Getenv("NODE1_NAME")
	dbPass := os.Getenv("NODE1_PASSWORD")

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	db.AutoMigrate(&model.Block{})

	return &databaseRepository{db}
}

// InitNode2 represents database configuration of Node2
func InitNode2() Repository {
	godotenv.Load()
	dbHost := os.Getenv("NODE2_HOST")
	dbPort := os.Getenv("NODE2_PORT")
	dbUser := os.Getenv("NODE2_USER")
	dbName := os.Getenv("NODE2_NAME")
	dbPass := os.Getenv("NODE2_PASSWORD")

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	db.AutoMigrate(&model.Block{})

	return &databaseRepository{db}
}

// InitNode3 represents database configuration of Node3
func InitNode3() Repository {
	godotenv.Load()
	dbHost := os.Getenv("NODE3_HOST")
	dbPort := os.Getenv("NODE3_PORT")
	dbUser := os.Getenv("NODE3_USER")
	dbName := os.Getenv("NODE3_NAME")
	dbPass := os.Getenv("NODE3_PASSWORD")

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	db.AutoMigrate(&model.Block{})

	return &databaseRepository{db}
}

// InitPool represents database configuration of Data Pool
func InitPool() Repository {
	godotenv.Load()
	dbHost := os.Getenv("POOL_HOST")
	dbPort := os.Getenv("POOL_PORT")
	dbUser := os.Getenv("POOL_USER")
	dbName := os.Getenv("POOL_NAME")
	dbPass := os.Getenv("POOL_PASSWORD")

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	db.AutoMigrate(&model.Akta{})

	return &databaseRepository{db}
}
