package dto

import (
	"github.com/AdiKhoironHasan/privy-cake-store/pkg/common/validator"
)

type CakeReqDTO struct {
	ID          int
	Title       string  `json:"title" valid:"required" validname:"title"`
	Description string  `json:"description" valid:"required" validname:"description"`
	Rating      float64 `json:"rating" valid:"required,maxlength:3" validname:"rating"`
	Image       string  `json:"image" valid:"required" validname:"image"`
}

func (dto *CakeReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type CakeParamReqDTO struct {
	Query  string `json:"query" validname:"query" query:"query"`
	Rating string `json:"rating" validname:"rating" query:"rating"`
}

func (dto *CakeParamReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}
