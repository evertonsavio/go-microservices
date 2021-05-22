package services

import (
	"github.com/havyx/golang-microservices/mvc/utils"
	"github.com/havyx/golang-microservices/mvc/domain"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError)  {
	return domain.GetUser(userId)
}