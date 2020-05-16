package service

import (
	"github.com/ehardi19/rantaiblok/model"
	"github.com/sirupsen/logrus"
)

// SaveAkta ...
func (s *Service) SaveAkta(akta model.Akta) error {

	err := s.Pool.SaveAkta(akta)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

// GetAllAkta ...
func (s *Service) GetAllAkta() ([]model.Akta, error) {
	aktas, err := s.Pool.GetAllAkta()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return aktas, err
}

// GetAktaByID ...
func (s *Service) GetAktaByID(id int) (model.Akta, error) {
	akta, err := s.Pool.GetAktaByID(id)
	if err != nil {
		logrus.Error(err)
		return model.Akta{}, err
	}

	return akta, err
}

// DeleteAktaByID ...
func (s *Service) DeleteAktaByID(id int) error {
	err := s.Pool.DeleteAktaByID(id)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return err
}
