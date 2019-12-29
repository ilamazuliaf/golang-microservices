package app

import (
	"github.com/ilamazuliaf/golang-microservices/mvc/controllers"
)

func mapUrls(){
	router.GET("/users/:user_id", controllers.GetUser)
}
