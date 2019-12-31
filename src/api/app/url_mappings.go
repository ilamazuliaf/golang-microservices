package app

import (
	"github.com/ilamazuliaf/golang-microservices/src/api/controllers/luffy"
	"github.com/ilamazuliaf/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/luffy", luffy.Luffy)
	router.POST("/repositories", repositories.CreateRepo)
}
