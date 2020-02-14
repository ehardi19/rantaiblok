package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ehardi19/rantaiblok/blockchain"
	"github.com/ehardi19/rantaiblok/blockchain/delivery/http"
	"github.com/ehardi19/rantaiblok/blockchain/repository"
	"github.com/ehardi19/rantaiblok/blockchain/usecase"
	"github.com/ehardi19/rantaiblok/middleware"
	"github.com/ehardi19/rantaiblok/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	godotenv.Load()
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASSWORD")
	port := os.Getenv("PORT")

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	repo := repository.NewGormRepository(db)
	usecase := usecase.NewUsecase(repo)

	err = initGenesis(repo)
	if err != nil {
		log.Fatal(err)
	}

	http.NewHandler(e, usecase)
	e.Logger.Fatal(e.Start(port))
}

func initGenesis(repo blockchain.Repository) error {
	check, _ := repo.Fetch()

	if len(check) > 0 {
		return nil
	}

	genesis := models.Block{
		ID:        0,
		Data:      "genesis",
		Timestamp: "",
		Hash:      "",
		PrevHash:  "",
	}
	err := repo.Store(genesis)
	if err != nil {
		return err
	}

	return nil
}
