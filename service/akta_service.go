package service

import (
	"github.com/ehardi19/rantaiblok/model"
	"github.com/sirupsen/logrus"
)

// SaveAkta ...
func (s *Service) SaveAkta(akta model.Akta) error {
	logrus.Info(akta)

	err := s.Repo.SaveAkta(akta)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

// GetAllAkta ...
func (s *Service) GetAllAkta() ([]model.Akta, error) {
	aktas, err := s.Repo.GetAllAkta()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	logrus.Info(aktas)

	return aktas, err
}

// GetAktaByID ...
func (s *Service) GetAktaByID(id int) (model.Akta, error) {
	akta, err := s.Repo.GetAktaByID(id)
	if err != nil {
		logrus.Error(err)
		return model.Akta{}, err
	}

	logrus.Info(akta)

	return akta, err
}

// GetAktaByAktaNum ...
func (s *Service) GetAktaByAktaNum(aktaNum string) (model.Akta, error) {
	akta, err := s.Repo.GetAktaByAktaNum(aktaNum)
	if err != nil {
		logrus.Error(err)
		return model.Akta{}, err
	}

	logrus.Info(akta)

	return akta, err
}
