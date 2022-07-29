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

func TestGetAllComment(t *testing.T) {
	repo := new(mocks.DataComment) // Menggunakan mock object yang sudah dibuat

	t.Run("Success get all", func(t *testing.T) {
		repo.On("GetAllComment").Return([]domain.Comment{{ID: 1, Comment: "pertamax gan", UserID: 1, ContentID: 1}}, nil).Once()
		// usecase := New(&mockUserDataTrue{})
		usecase := New(repo, validator.New())
		res, err := usecase.GetAllComment()
		assert.Nil(t, err)
		assert.GreaterOrEqual(t, len(res), 1)
		assert.Greater(t, res[0].ID, 0)
		repo.AssertExpectations(t)
	})

	t.Run("Error not found", func(t *testing.T) {
		repo.On("GetAllComment").Return(nil, gorm.ErrRecordNotFound).Once()
		// usecase := New(&mockUserDataFalse{})
		usecase := New(repo, validator.New())
		res, err := usecase.GetAllComment()
		assert.NotNil(t, err)
		assert.EqualError(t, err, gorm.ErrRecordNotFound.Error())
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error cannot retrieve data", func(t *testing.T) {
		repo.On("GetAllComment").Return(nil, errors.New("cannot retrieve data")).Once()
		// usecase := New(&mockUserDataFalse{})
		usecase := New(repo, validator.New())
		res, err := usecase.GetAllComment()
		assert.NotNil(t, err)
		assert.EqualError(t, err, errors.New("error when retrieve data").Error())
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}

func TestInsertComment(t *testing.T) {
	repo := new(mocks.DataComment)

	mockData := domain.Comment{Comment: "Pertamax Gan"}
	returnData := mockData
	returnData.UserID = 1
	returnData.ContentID = 1
	returnData.ID = 1
	t.Run("Success post content", func(t *testing.T) {
		repo.On("PostComment", mock.Anything).Return(returnData, nil).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.PostingComment(returnData.ContentID, mockData)
		assert.Nil(t, err)
		assert.Greater(t, res.ID, 0)
		assert.Equal(t, "Pertamax Gan", res.Comment)
		repo.AssertExpectations(t)
	})
}
