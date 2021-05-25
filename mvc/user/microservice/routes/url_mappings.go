package routes

import(
	// "net/http"
	"github.com/havyx/golang-microservices/mvc/user/microservice/handlers"
)

func mapUrls(){

	router.GET("/users/:user_id", handlers.GetUser)
	// http.HandleFunc("/users", handlers.GetUser)
}