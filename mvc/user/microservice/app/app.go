package app

import (
	"github.com/havyx/golang-microservices/mvc/user/microservice/controllers"
	"net/http"
)

func StartApp()  {
	http.HandleFunc("/users", controllers.GetUser)
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}