package assembler

import (
	"github.com/AdiKhoironHasan/privy-cake-store/internal/models"
	"github.com/AdiKhoironHasan/privy-cake-store/pkg/dto"
)

func ToAddNewCake(d *dto.CakeReqDTO) *models.CakeModels {
	return &models.CakeModels{
		Title:       d.Title,
		Description: d.Description,
		Rating:      d.Rating,
		Image:       d.Image,
	}
}

func ToUpdateCake(d *dto.CakeReqDTO) *models.CakeModels {
	return &models.CakeModels{
		ID:          d.ID,
		Title:       d.Title,
		Description: d.Description,
		Rating:      d.Rating,
		Image:       d.Image,
	}
}

func ToListOfCakeResponse(d []*models.CakeModels) []*dto.CakeResDTO {
	var dataCakes []*dto.CakeResDTO

	for _, val := range d {
		dataCakes = append(dataCakes, &dto.CakeResDTO{
			ID:    val.ID,
			Title: val.Title,
			// Description: val.Description,
			Rating: val.Rating,
			Image:  val.Image,
			// Created_at:  val.Created_at,
			// Updated_at:  val.Updated_at,
		})
	}

	return dataCakes
}

func ToShowCakeByIDResponse(d []*models.CakeModels) *dto.CakeResDTO {
	var dataCake *dto.CakeResDTO
	dataCake = &dto.CakeResDTO{
		ID:          d[0].ID,
		Title:       d[0].Title,
		Description: d[0].Description,
		Rating:      d[0].Rating,
		Image:       d[0].Image,
		Created_at:  d[0].Created_at,
		Updated_at:  d[0].Updated_at,
	}

	return dataCake
}
