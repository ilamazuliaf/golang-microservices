package luffy

import (
	"github.com/ilamazuliaf/golang-microservices/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t,"D Monkey", luffy)
}

func TestLuffy(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet,"/luffy", nil)
	c := test_utils.GetMockedContext(request, response)

	Luffy(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "D Monkey", response.Body.String())
}