package services

import (
	integ "github.com/AdiKhoironHasan/privy-cake-store/internal/integration"
	"github.com/AdiKhoironHasan/privy-cake-store/internal/repository"
	"github.com/AdiKhoironHasan/privy-cake-store/pkg/dto"
	"github.com/AdiKhoironHasan/privy-cake-store/pkg/dto/assembler"
)

type service struct {
	sqlRepo repository.SqlRepository
	// noSqlRepo repository.NoSqlRepository
	IntegServ integ.IntegServices
}

func NewService(sqlRepo repository.SqlRepository, IntegServ integ.IntegServices) Services {
	return &service{sqlRepo, IntegServ}
}

func (s *service) AddNewCake(req *dto.CakeReqDTO) error {
	dataCakeModel := assembler.ToAddNewCake(req)

	err := s.sqlRepo.AddNewCake(dataCakeModel)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) ShowAllCake(req *dto.CakeParamReqDTO) ([]*dto.CakeResDTO, error) {
	var dataCakes []*dto.CakeResDTO
	var where string

	if req.Query != "" && req.Rating != "" {
		where = "title LIKE '%" + req.Query + "%' AND description LIKE '%" + req.Query + "%' AND rating = '" + req.Rating + "'"
	} else if req.Query != "" {
		where = "title LIKE '%" + req.Query + "%' AND description LIKE '%" + req.Query + "%'"
	} else if req.Rating != "" {
		where = "rating = '" + req.Rating + "'"
	}

	dataCakeModels, err := s.sqlRepo.ShowAllCake(where)
	if err != nil {
		return nil, err
	}

	dataCakes = assembler.ToListOfCakeResponse(dataCakeModels)

	return dataCakes, nil
}

func (s *service) ShowCakeByID(id int) (*dto.CakeResDTO, error) {
	var dataCake *dto.CakeResDTO

	dataCakeModels, err := s.sqlRepo.ShowCakeByID(id)
	if err != nil {
		return nil, err
	}

	if dataCakeModels == nil {
		return nil, nil
	}

	dataCake = assembler.ToShowCakeByIDResponse(dataCakeModels)

	return dataCake, nil
}

func (s *service) UpdateCake(req *dto.CakeReqDTO) error {
	dataCake := assembler.ToUpdateCake(req)

	err := s.sqlRepo.UpdateCake(dataCake)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteCake(id int) error {

	err := s.sqlRepo.DeleteCake(id)
	if err != nil {
		return err
	}

	return nil
}
