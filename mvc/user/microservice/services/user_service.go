package services

import (
	"github.com/havyx/golang-microservices/mvc/user/microservice/domain"
	"github.com/havyx/golang-microservices/mvc/user/microservice/utils"
)

type userService struct {
}

var (
	UserService userService
)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
