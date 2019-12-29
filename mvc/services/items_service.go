package services

import (
	"github.com/ilamazuliaf/golang-microservices/mvc/domain"
	"github.com/ilamazuliaf/golang-microservices/mvc/utils"
	"net/http"
)

type itemService struct {

}

var (
	ItemService itemService
)

func (u *itemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return  nil, &utils.ApplicationError{
		Message:    "Implement me",
		StatusCode: http.StatusInternalServerError,
		Code:       "internal server error",
	}
}