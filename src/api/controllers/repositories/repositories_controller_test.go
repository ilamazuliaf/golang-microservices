package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/ilamazuliaf/golang-microservices/src/api/utils/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateRepoInvalidJasonRequest(t *testing.T)  {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))
	c.Request = request

	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid json body", apiErr.Message())
}

func TestCreateRepoErrorFromGithub(t *testing.T)  {
	// TODO: next time
}