package routes

import (
	"net/http"

	"github.com/havyx/golang-microservices/mvc/user/microservice/handlers"
)

func StartApp() {
	http.HandleFunc("/users", handlers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
