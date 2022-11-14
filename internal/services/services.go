package services

import "github.com/AdiKhoironHasan/privy-cake-store/pkg/dto"

type Services interface {
	AddNewCake(req *dto.CakeReqDTO) error
	ShowAllCake(req *dto.CakeParamReqDTO) ([]*dto.CakeResDTO, error)
	ShowCakeByID(id int) (*dto.CakeResDTO, error)
	UpdateCake(req *dto.CakeReqDTO) error
	DeleteCake(id int) error
}
