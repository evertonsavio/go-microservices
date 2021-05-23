package routes

import (
	"github.com/havyx/golang-microservices/mvc/user/microservice/handlers"
	"net/http"
)

func StartApp()  {
	http.HandleFunc("/users", handlers.GetUser)
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}