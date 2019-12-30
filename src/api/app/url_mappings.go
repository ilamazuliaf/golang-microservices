package app

import (
	"github.com/ilamazuliaf/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.POST("/repositories", repositories.CreateRepo)
}
