package services

import (
	"errors"
	"testing"

	"github.com/AdiKhoironHasan/privy-cake-store/internal/models"
	"github.com/AdiKhoironHasan/privy-cake-store/internal/repository"
	"github.com/AdiKhoironHasan/privy-cake-store/pkg/dto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repoMock = &repository.RepositoryMock{Mock: mock.Mock{}}
var servMock = service{sqlRepo: repoMock}

func TestAddNewCake_Success(t *testing.T) {
	models := models.CakeModels{
		Title:       "Mango cheesecake",
		Description: "A cheesecake made of mango",
		Rating:      7.5,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	}

	cake := dto.CakeReqDTO{
		Title:       "Mango cheesecake",
		Description: "A cheesecake made of mango",
		Rating:      7.5,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	}

	repoMock.Mock.On("AddNewCake", &models).Return(nil)

	err := servMock.AddNewCake(&cake)
	assert.Nil(t, err)
}

func TestAddNewCake_Error(t *testing.T) {
	models := models.CakeModels{
		Title:       "Mango cheesecake 2",
		Description: "A cheesecake made of mango 2",
		Rating:      7,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	}

	cake := dto.CakeReqDTO{
		Title:       "Mango cheesecake 2",
		Description: "A cheesecake made of mango 2",
		Rating:      7,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	}

	repoMock.Mock.On("AddNewCake", &models).Return(errors.New(mock.Anything))

	err := servMock.AddNewCake(&cake)
	assert.NotNil(t, err)
}

func TestShowAllCake_Success(t *testing.T) {
	models := []*models.CakeModels{
		{
			ID:          1,
			Title:       "Mango cheesecake",
			Description: "A cheesecake made of mango",
			Rating:      7.5,
			Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
			Created_at:  "2022-11-14 19:56:22",
			Updated_at:  "2022-11-14 19:56:22",
		},
	}

	response := []*dto.CakeResDTO{
		{
			ID:     1,
			Title:  "Mango cheesecake",
			Rating: 7.5,
			Image:  "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
		},
	}

	where := ""

	repoMock.Mock.On("ShowAllCake", where).Return(models)

	cakeParam := dto.CakeParamReqDTO{
		Query:  "",
		Rating: "",
	}
	cakes, err := servMock.ShowAllCake(&cakeParam)
	assert.Nil(t, err)
	assert.NotNil(t, cakes)
	assert.Equal(t, response, cakes)
}

func TestShowAllCake_Error(t *testing.T) {
	where := "title LIKE '%cheese%' AND description LIKE '%cheese%' AND rating = '7'"

	repoMock.Mock.On("ShowAllCake", where).Return(nil)

	cakeParam := dto.CakeParamReqDTO{
		Query:  "cheese",
		Rating: "7",
	}
	cakes, err := servMock.ShowAllCake(&cakeParam)
	assert.Nil(t, cakes)
	assert.NotNil(t, err)
}

func TestShowAllCake_ErrorQuery(t *testing.T) {
	where := "title LIKE '%cheese%' AND description LIKE '%cheese%'"

	repoMock.Mock.On("ShowAllCake", where).Return(nil)

	cakeParam := dto.CakeParamReqDTO{
		Query:  "cheese",
		Rating: "",
	}
	cakes, err := servMock.ShowAllCake(&cakeParam)
	assert.Nil(t, cakes)
	assert.NotNil(t, err)
}

func TestShowAllCake_ErrorRating(t *testing.T) {
	where := "rating = '7'"

	repoMock.Mock.On("ShowAllCake", where).Return(nil)

	cakeParam := dto.CakeParamReqDTO{
		Query:  "",
		Rating: "7",
	}
	cakes, err := servMock.ShowAllCake(&cakeParam)
	assert.Nil(t, cakes)
	assert.NotNil(t, err)
}

func TestShowCakeByID_NotFound(t *testing.T) {
	repoMock.Mock.On("ShowCakeByID", 1).Return(nil)

	cake, err := servMock.ShowCakeByID(1)
	assert.Nil(t, cake)
	assert.NotNil(t, err)
}

func TestShowCakeByID_Success(t *testing.T) {
	models := []*models.CakeModels{
		{
			ID:          2,
			Title:       "Mango cheesecake",
			Description: "A cheesecake made of mango",
			Rating:      7.5,
			Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
			Created_at:  "2022-11-14 19:56:22",
			Updated_at:  "2022-11-14 19:56:22",
		},
	}
	repoMock.Mock.On("ShowCakeByID", 2).Return(models)

	cake, err := servMock.ShowCakeByID(2)
	assert.Nil(t, err)
	assert.NotNil(t, cake)
	assert.Equal(t, models[0].ID, cake.ID)
	assert.Equal(t, models[0].Title, cake.Title)
	assert.Equal(t, models[0].Description, cake.Description)
	assert.Equal(t, models[0].Rating, cake.Rating)
	assert.Equal(t, models[0].Image, cake.Image)
	assert.Equal(t, models[0].Created_at, cake.Created_at)
	assert.Equal(t, models[0].Updated_at, cake.Updated_at)
}

func TestUpdateCake_Success(t *testing.T) {
	models := models.CakeModels{
		Title:       "Mango cheesecake",
		Description: "A cheesecake made of mango",
		Rating:      7.5,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	}

	cake := dto.CakeReqDTO{
		Title:       "Mango cheesecake",
		Description: "A cheesecake made of mango",
		Rating:      7.5,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	}

	repoMock.Mock.On("UpdateCake", &models).Return(nil)

	err := servMock.UpdateCake(&cake)
	assert.Nil(t, err)
}

func TestUpdateCake_Error(t *testing.T) {
	models := models.CakeModels{
		Title:       "Mango cheesecake 2",
		Description: "A cheesecake made of mango 2",
		Rating:      7,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	}

	cake := dto.CakeReqDTO{
		Title:       "Mango cheesecake 2",
		Description: "A cheesecake made of mango 2",
		Rating:      7,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	}

	repoMock.Mock.On("UpdateCake", &models).Return(errors.New(mock.Anything))

	err := servMock.UpdateCake(&cake)
	assert.NotNil(t, err)
}

func TestDeleteCake_Success(t *testing.T) {
	repoMock.Mock.On("DeleteCake", 1).Return(nil)

	err := servMock.DeleteCake(1)
	assert.Nil(t, err)
}

func TestDeleteCake_Error(t *testing.T) {
	repoMock.Mock.On("DeleteCake", 2).Return(errors.New(mock.Anything))

	err := servMock.DeleteCake(2)
	assert.NotNil(t, err)
}
