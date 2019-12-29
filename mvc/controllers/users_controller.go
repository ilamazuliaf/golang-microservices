package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilamazuliaf/golang-microservices/mvc/services"
	"github.com/ilamazuliaf/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context){
	// Get Query Url
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user id must be number",
			StatusCode: http.StatusBadRequest,
			Code:       "Bad Request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, user)
	//c.JSON(http.StatusOK, user)
}