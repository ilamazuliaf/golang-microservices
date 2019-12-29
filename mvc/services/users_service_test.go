package services

import (
	"github.com/ilamazuliaf/golang-microservices/mvc/domain"
	"github.com/ilamazuliaf/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	userDaoMock usersDaoMock
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

type usersDaoMock struct {

}

func init() {
	domain.UserDao = &usersDaoMock{}
}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError){
	return  getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return  nil, &utils.ApplicationError{
			Message:    "user 0 does not exists",
			StatusCode: http.StatusNotFound,
		}
	}
	user, err := UsersService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t,"user 0 does not exists", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id:        123,
			FirstName: "Faizul",
			LastName:  "Amaly",
			Email:     "faiz@gmail.com",
		}, nil
	}
	user, err := UsersService.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t,123, user.Id)
}