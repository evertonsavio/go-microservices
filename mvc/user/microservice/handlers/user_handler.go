package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/havyx/golang-microservices/mvc/user/microservice/services"
	"github.com/havyx/golang-microservices/mvc/user/microservice/utils"
)

func GetUser(c *gin.Context){//(resp http.ResponseWriter, req *http.Request) {

	//userId := req.URL.Query().Get("user_id")
	userId := c.Param("user_id")

	log.Println("UserId is: " + userId)

	userIdInt, err := strconv.ParseInt(userId, 10, 64)

	if err != nil {

		apiErr := &utils.ApplicationError{
			Message:    "UserId must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       400,
		}
		utils.Respond(c, apiErr.StatusCode, apiErr)
		//c.JSON(http.StatusBadRequest, apiErr)
		return
		// jsonValue, _ := json.Marshal(apiErr)
		// resp.WriteHeader(apiErr.StatusCode)
		// resp.Write(jsonValue)
	}

	user, apiErr := services.UserService.GetUser(userIdInt)
	if apiErr != nil {
		// jsonValue, _ := json.Marshal(apiErr)
		// resp.WriteHeader(apiErr.StatusCode)
		// resp.Write(jsonValue)
		utils.Respond(c, apiErr.StatusCode, apiErr)
		//c.JSON(apiErr.StatusCode, apiErr)
		return
	}
	//curl localhost:8080/users/123 -H "Accept:application/xml" -v
	utils.Respond(c, http.StatusOK, user)
	//c.JSON(http.StatusOK, user)
	// jsonValue, _ := json.Marshal(user)
	// resp.Write(jsonValue)
}