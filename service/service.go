package service

import (
	"github.com/ehardi19/rantaiblok/repository"
)

// Service ...
type Service struct {
	Repo repository.Repository
}

// New ...
func New() Service {
	return Service{
		Repo: repository.Init(),
	}
}
