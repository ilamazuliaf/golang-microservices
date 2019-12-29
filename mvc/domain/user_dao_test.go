package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	// inisialisasi

	// eksekusi
	user ,err := GetUser(0)

	// validasi
	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "Not Found", err.Code)
	assert.EqualValues(t, "user 0 does not exists",err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Faizul", user.FirstName)
	assert.EqualValues(t, "Amaly", user.LastName)
	assert.EqualValues(t, "faiz@gmail.com", user.Email)
}