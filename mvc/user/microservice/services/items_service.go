package services

import (
	"github.com/havyx/golang-microservices/mvc/user/microservice/domain"
	"github.com/havyx/golang-microservices/mvc/user/microservice/utils"
	"net/http"
)

type itemService struct {
}

var (
	ItemService itemService
)

func (s *itemService) GetItem(itemId string)(*domain.Item, *utils.ApplicationError){
	return nil, &utils.ApplicationError{
		Message: "Implemented me",
		StatusCode: http.StatusInternalServerError,
	}
}