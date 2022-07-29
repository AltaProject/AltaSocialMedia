package usecase

import (
	"errors"
	"testing"

	"github.com/AltaProject/AltaSocialMedia/domain"
	"github.com/AltaProject/AltaSocialMedia/domain/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestGetAllContent(t *testing.T) {
	repo := new(mocks.ContentData) // Menggunakan mock object yang sudah dibuat

	t.Run("Success get all", func(t *testing.T) {
		repo.On("GetAllContent").Return([]domain.Content{{ID: 1, Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", UserID: 1}}, nil).Once()
		// usecase := New(&mockUserDataTrue{})
		usecase := New(repo, validator.New())
		res, err := usecase.GetAllContent()
		assert.Nil(t, err)
		assert.GreaterOrEqual(t, len(res), 1)
		assert.Greater(t, res[0].ID, 0)
		repo.AssertExpectations(t)
	})

	t.Run("Error not found", func(t *testing.T) {
		repo.On("GetAllContent").Return(nil, gorm.ErrRecordNotFound).Once()
		// usecase := New(&mockUserDataFalse{})
		usecase := New(repo, validator.New())
		res, err := usecase.GetAllContent()
		assert.NotNil(t, err)
		assert.EqualError(t, err, gorm.ErrRecordNotFound.Error())
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error cannot retrieve data", func(t *testing.T) {
		repo.On("GetAllContent").Return(nil, errors.New("cannot retrieve data")).Once()
		// usecase := New(&mockUserDataFalse{})
		usecase := New(repo, validator.New())
		res, err := usecase.GetAllContent()
		assert.NotNil(t, err)
		assert.EqualError(t, err, errors.New("error when retrieve data").Error())
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}

func TestPostContent(t *testing.T) {
	repo := new(mocks.ContentData)

	mockData := domain.Content{Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit."}
	returnData := mockData
	returnData.UserID = 1
	returnData.ID = 1
	t.Run("Success post content", func(t *testing.T) {
		repo.On("AddNewContent", mock.Anything).Return(returnData, nil).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.Posting(returnData.UserID, mockData)
		assert.Nil(t, err)
		assert.Greater(t, res.ID, 0)
		assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", res.Content)
		repo.AssertExpectations(t)
	})
}
