package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRestError(t *testing.T) {
	err := NewRestError("this is the message", http.StatusNotImplemented, "not_implemented", nil)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotImplemented, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "message: this is the message - status: 501 - error: not_implemented - causes: []", err.Error())

	assert.Nil(t, err.Causes())
}

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "message: this is the message - status: 500 - error: internal_server_error - causes: [database error]", err.Error())

	assert.NotNil(t, err.Causes())
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "database error", err.Causes()[0])
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "message: this is the message - status: 400 - error: bad_request - causes: []", err.Error())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "message: this is the message - status: 404 - error: not_found - causes: []", err.Error())
}

func TestNewUnauthorizedError(t *testing.T) {
	err := NewUnauthorizedError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "message: this is the message - status: 401 - error: unauthorized - causes: []", err.Error())
}

func TestNewRestErrorFromBytes(t *testing.T) {
	restErr, err := NewRestErrorFromBytes([]byte("some"))
	assert.Nil(t, restErr)
	assert.NotNil(t, err)
}

func TestNewError(t *testing.T) {
	err := NewError("this is the message")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "this is the message")
}
