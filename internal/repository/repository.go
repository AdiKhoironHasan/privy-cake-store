package repository

import (
	"github.com/AdiKhoironHasan/privy-cake-store/internal/models"
)

type SqlRepository interface {
	AddNewCake(dataCake *models.CakeModels) error
	ShowAllCake(where string) ([]*models.CakeModels, error)
	ShowCakeByID(id int) ([]*models.CakeModels, error)
	UpdateCake(dataCake *models.CakeModels) error
	DeleteCake(id int) error
}
