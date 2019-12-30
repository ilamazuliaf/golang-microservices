package services

import (
	"github.com/ilamazuliaf/golang-microservices/src/api/config"
	"github.com/ilamazuliaf/golang-microservices/src/api/domain/github"
	"github.com/ilamazuliaf/golang-microservices/src/api/domain/repositories"
	"github.com/ilamazuliaf/golang-microservices/src/api/providers/github_provider"
	"github.com/ilamazuliaf/golang-microservices/src/api/utils/errors"
	"strings"
)

type reposService struct {
}

type reposServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Owner: response.Owner.Login,
		Name:  response.Name,
	}
	return &result, nil
}
