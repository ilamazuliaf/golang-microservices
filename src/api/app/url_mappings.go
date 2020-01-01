package app

import (
	"github.com/ilamazuliaf/golang-microservices/src/api/controllers/luffy"
	"github.com/ilamazuliaf/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/luffy", luffy.Luffy)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
