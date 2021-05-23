package services

import (
	"github.com/havyx/golang-microservices/mvc/user/microservice/utils"
	"github.com/havyx/golang-microservices/mvc/user/microservice/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"net/http"
)

var(
	userDaoMock usersDaoMock

	getUserFunc func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init()  {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}


func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunc(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunc = func(userId int64)(*domain.User, *utils.ApplicationError){
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message: "User 0 was not found",
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "User 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunc = func(userId int64)(*domain.User, *utils.ApplicationError){
		return &domain.User{
			Id: 123,
		}, nil
	}
	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}