package main

import (
	"log"

	"github.com/ehardi19/rantaiblok/handler"
	"github.com/ehardi19/rantaiblok/model"
	"github.com/ehardi19/rantaiblok/repository"
	"github.com/ehardi19/rantaiblok/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func main() {

	s := service.New()
	h := handler.InitHandler(s)
	e := echo.New()

	err := initGenesis(s.Node1, s.Node2, s.Node3)
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

func initGenesis(node1, node2, node3 repository.Repository) error {
	check, _ := node1.GetAllBlock()

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

	// Saving to Node1
	err := node1.SaveBlock(genesis)
	if err != nil {
		return err
	}

	// Saving to Node2
	err = node2.SaveBlock(genesis)
	if err != nil {
		return err
	}

	// Saving to Node3
	err = node3.SaveBlock(genesis)
	if err != nil {
		return err
	}

	logrus.Println("genesis created")

	return nil
}
