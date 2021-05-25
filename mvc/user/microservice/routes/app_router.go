package routes

import (
	// "net/http"

	"github.com/gin-gonic/gin"
)

var(
	router *gin.Engine
)


func init(){
	router = gin.Default()
}

func StartApp() {

	mapUrls()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	panic(err)
	// }
}
