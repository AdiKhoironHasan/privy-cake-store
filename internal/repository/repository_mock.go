package repository

import (
	"errors"

	"github.com/AdiKhoironHasan/privy-cake-store/internal/models"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	Mock mock.Mock
}

func (repo *RepositoryMock) AddNewCake(dataCake *models.CakeModels) error {
	arguments := repo.Mock.Called(dataCake)
	var err error
	if n, ok := arguments.Get(0).(error); ok {
		err = n
	}

	return err
}

func (repo *RepositoryMock) ShowAllCake(where string) ([]*models.CakeModels, error) {
	arguments := repo.Mock.Called(where)
	if arguments.Get(0) == nil {
		return nil, errors.New(mock.Anything)
	} else {
		cake := arguments.Get(0).([]*models.CakeModels)
		return cake, nil
	}
}
func (repo *RepositoryMock) ShowCakeByID(id int) ([]*models.CakeModels, error) {
	arguments := repo.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil, errors.New(mock.Anything)
	} else {
		cake := arguments.Get(0).([]*models.CakeModels)
		return cake, nil
	}
}
func (repo *RepositoryMock) UpdateCake(dataCake *models.CakeModels) error {
	arguments := repo.Mock.Called(dataCake)
	var err error
	if n, ok := arguments.Get(0).(error); ok {
		err = n
	}

	return err
}
func (repo *RepositoryMock) DeleteCake(id int) error {
	arguments := repo.Mock.Called(id)
	var err error
	if n, ok := arguments.Get(0).(error); ok {
		err = n
	}

	return err
}
