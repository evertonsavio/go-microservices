package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsetNotFound(t *testing.T) {
	user, err := GetUser(0)
	if user != nil {
		t.Error("We were not expecting user with id 0.")
	}
	if err == nil {
		t.Error("We were expecting an error when user id is 0.")
	}
	if err.StatusCode != http.StatusNotFound {
		t.Error("We were expecting 404 when user not found.")
	}
}

func TestGetUsetNotFoundWithAssertPackage(t *testing.T) {

	user, err := GetUser(0)

	assert.Nil(t, user, "We were not expecting user with id 0.")
	assert.NotNil(t, err, "We were expecting an error when user id is 0.")
	assert.EqualValues(t, 404, err.Code, "We were expecting an error code 404")
	assert.EqualValues(t,http.StatusNotFound, err.StatusCode, "We were expecting 404 when user not found." )

}

func TestGetUserNoError(t *testing.T){

	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Savio", user.FistName)
	assert.EqualValues(t, "Lucas", user.LastName)
	assert.EqualValues(t, "everluca@hotmail.com", user.Email)
	
}

