package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/havyx/golang-microservices/mvc/user/microservice/utils"
)

var (
	users = map[int64]*User{
		123: &User{Id: 123, FistName: "Savio", LastName: "Lucas", Email: "everluca@hotmail.com"},
	}

	//UserDao userDao
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {

	log.Println("Access DB")

	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       404,
	}
}
