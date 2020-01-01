package repositories

import (
	"github.com/ilamazuliaf/golang-microservices/src/api/utils/errors"
	"strings"
)

type CreateRepoRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

// Create validate function repo name
func (r *CreateRepoRequest) Validate() errors.ApiError {
	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		return  errors.NewBadRequestError("invalid respository name")
	}
	return nil
}

type CreateRepoResponse struct {
	Id    int64  `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type CreateReposResponse struct {
	StatusCode int `json:"status"`
	Results []CreateRepositoriesResult `json:"results"`
}

type CreateRepositoriesResult struct {
	Response CreateRepoResponse `json:"respo"`
	Error errors.ApiError `json:"error"`
}