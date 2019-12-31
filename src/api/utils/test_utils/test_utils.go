package test_utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func GetMockedContext(request *http.Request, response *httptest.ResponseRecorder) *gin.Context {
	c, _ := gin.CreateTestContext(response)
	c.Request = request
	return c
}
