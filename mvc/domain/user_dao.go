package domain

import (
	"fmt"
	"github.com/ilamazuliaf/golang-microservices/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123:&User{Id:123, FirstName:"Faizul", LastName:"Amaly", Email:"faiz@gmail.com"},
	}

	UserDao userDaoInterface
)

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {

}

func init() {
	UserDao = &userDao{}
}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we're accessing the database")
	if user := users[userId]; user != nil {
		return  user, nil
	}

	return  nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exists", userId),
		StatusCode: http.StatusNotFound,
		Code:       "Not Found",
	}
}