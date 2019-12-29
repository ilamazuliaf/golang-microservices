package services

import (
	"github.com/ilamazuliaf/golang-microservices/mvc/domain"
	"github.com/ilamazuliaf/golang-microservices/mvc/utils"
)

type userServices struct {

}

var (
	UsersService userServices
)

func (u *userServices) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	//return domain.UserDao.GetUser(userId)
	user, err := domain.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return  user, nil
}
