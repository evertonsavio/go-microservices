package domain

import (
	"fmt"
	"net/http"

	"github.com/havyx/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: &User{Id: 123, FistName: "Savio", LastName: "Lucas", Email: "everluca@hotmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       404,
	}
}
