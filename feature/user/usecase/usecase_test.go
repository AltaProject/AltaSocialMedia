package usecase

import (
	"testing"

	"github.com/AltaProject/AltaSocialMedia/domain"
	"github.com/AltaProject/AltaSocialMedia/domain/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestRegister(t *testing.T) {
	repo := new(mocks.UserData)

	mockData := domain.User{Nama: "Jon", Username: "Jondoe", Email: "jondoe@mail.com", Password: "jon123", No_HP: "0216837618"}
	returnData := mockData
	returnData.ID = 1
	returnData.Password = "$2a$10$OqHN2OI/X2g8c5on5JV33.m0vLv4U5nhniXpb.hu2ddcSSj/nZMFq"
	t.Run("Success register", func(t *testing.T) {
		repo.On("Register", mock.Anything).Return(returnData, nil).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.Register(mockData)
		assert.Nil(t, err)
		assert.Greater(t, res.ID, 0)
		assert.Equal(t, "Jon", res.Nama)
		assert.Equal(t, "Jondoe", res.Username)
		assert.Equal(t, "jondoe@mail.com", res.Email)
		assert.Equal(t, "$2a$10$OqHN2OI/X2g8c5on5JV33.m0vLv4U5nhniXpb.hu2ddcSSj/nZMFq", res.Password, "Password tidak sesuai")
		assert.Equal(t, "0216837618", res.No_HP)
		repo.AssertExpectations(t)
	})
	t.Run("Duplicated Data", func(t *testing.T) {
		repo.On("Register", mock.Anything).Return(domain.User{}, gorm.ErrRegistered).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.Register(returnData)
		assert.NotNil(t, err)
		assert.EqualError(t, err, gorm.ErrRegistered.Error())
		assert.Equal(t, 0, res.ID)
		assert.Equal(t, "", res.Nama)
		assert.Equal(t, "", res.Username)
		assert.Equal(t, "", res.Email)
		assert.Equal(t, "", res.Password)
		assert.Equal(t, "", res.No_HP)
		repo.AssertExpectations(t)
	})

	// t.Run("Validator error", func(t *testing.T) {
	//  // useCase := New(&mockUserDataTrue{})
	//  repo.On("Register", mock.Anything).Return(returnData, nil).Once()
	//  useCase := New(repo, validator.New())
	//  res, err := useCase.Register(domain.User{})
	//  assert.EqualError(t, err, "error") // Apakah errornya nil
	//  assert.Greater(t, res.ID, 0)       // Apakah ID nya lebih besar dari 0
	//  assert.Equal(t, "", res.Nama)      // Apakah nama yang di insertkan sama
	//  assert.Equal(t, "", res.Email)
	//  assert.Equal(t, "", res.Password, "Password tidak sesuai")
	//  assert.Equal(t, "", res.No_HP)
	//  repo.AssertExpectations(t)
	// })
}

func TestGetUserById(t *testing.T) {
	repo := new(mocks.UserData)

	mockData := domain.User{Nama: "Jon", Username: "Jondoe", Email: "jondoe@mail.com", Password: "jon123", No_HP: "0216837618"}
	returnData := mockData
	returnData.ID = 1
	returnData.Password = "$2a$10$OqHN2OI/X2g8c5on5JV33.m0vLv4U5nhniXpb.hu2ddcSSj/nZMFq"
	t.Run("Success Get Data By Id", func(t *testing.T) {
		repo.On("GetSpecificUser", mock.Anything).Return(returnData, nil).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.GetSpecificUser(1)
		assert.Nil(t, err)
		assert.Greater(t, res.ID, 0)
		assert.Equal(t, "Jon", res.Nama)
		assert.Equal(t, "Jondoe", res.Username)
		assert.Equal(t, "jondoe@mail.com", res.Email)
		assert.Equal(t, "$2a$10$OqHN2OI/X2g8c5on5JV33.m0vLv4U5nhniXpb.hu2ddcSSj/nZMFq", res.Password, "Password tidak sesuai")
		assert.Equal(t, "0216837618", res.No_HP)
		repo.AssertExpectations(t)
	})

	// t.Run("Error not found", func(t *testing.T) {
	//  repo.On("GetUserById").Return(nil, gorm.ErrRecordNotFound).Once()
	//  usecase := New(&mockUserDataFalse{})
	//  // usecase := New(repo, validator.New())
	//  res, err := usecase.GetSpecificUser(0)
	//  assert.NotNil(t, err)
	//  assert.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	//  assert.Nil(t, res)
	//  repo.AssertExpectations(t)
	// })
}
