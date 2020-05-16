package main

import (
	"log"

	"github.com/ehardi19/rantaiblok/handler"
	"github.com/ehardi19/rantaiblok/model"
	"github.com/ehardi19/rantaiblok/repository"
	"github.com/ehardi19/rantaiblok/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	s := service.New()
	h := handler.InitHandler(s)
	e := echo.New()

	err := initGenesis(s.Repo)
	if err != nil {
		log.Fatal(err)
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	e.GET("/", h.HelloWorld)
	e.GET("/block", h.GetAllBlock)
	e.GET("/block/:id", h.GetBlockByID)
	e.GET("/block/last", h.GetLastBlock)
	e.POST("/block", h.SaveBlock)
	e.GET("/valid", h.IsValid)
	e.Logger.Fatal(e.Start(":8000"))
}

func initGenesis(repo repository.Repository) error {
	check, _ := repo.GetAllBlock()

	if len(check) > 0 {
		return nil
	}

	genesis := model.Block{
		ID:        0,
		Data:      "genesis",
		Timestamp: "",
		Hash:      "",
		PrevHash:  "",
	}
	err := repo.SaveBlock(genesis)
	if err != nil {
		return err
	}

	return nil
}
