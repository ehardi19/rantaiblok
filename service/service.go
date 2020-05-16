package service

import (
	"github.com/ehardi19/rantaiblok/repository"
)

// Service defines how service build and all available services
type Service struct {
	Node1 repository.Repository
	Node2 repository.Repository
	Node3 repository.Repository
	Pool  repository.Repository
}

// New creates new services using nodes and data pool
func New() Service {
	return Service{
		Node1: repository.InitNode1(),
		Node2: repository.InitNode2(),
		Node3: repository.InitNode3(),
		Pool:  repository.InitPool(),
	}
}
