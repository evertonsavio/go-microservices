package services

import (
	"github.com/havyx/golang-microservices/mvc/user/microservice/utils"
	"github.com/havyx/golang-microservices/mvc/user/microservice/domain"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError)  {
	return domain.GetUser(userId)
}