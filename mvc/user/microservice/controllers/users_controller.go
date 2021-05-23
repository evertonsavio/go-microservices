package controllers

import (
	"github.com/havyx/golang-microservices/mvc/user/microservice/utils"
	"encoding/json"
	"github.com/havyx/golang-microservices/mvc/user/microservice/services"
	"strconv"
	"log"
	"net/http"
)

func GetUser(resp http.ResponseWriter, req *http.Request){

	userId := req.URL.Query().Get("user_id")
	log.Println("UserId is: "+ userId)

	userIdInt, err := strconv.ParseInt(userId, 10, 64)

	if err != nil {

		apiErr := &utils.ApplicationError{
			Message:    "UserId must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       400,
		}

		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	
	user,apiErr := services.GetUser(userIdInt)

	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}