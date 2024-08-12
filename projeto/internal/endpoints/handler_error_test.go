package endpoints

import (
	localerrors "emailn/internal/local-errors"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandlerError_InternalServerError(t *testing.T) {
	assert := assert.New(t)
	endpointFunc := func(_ http.ResponseWriter, _ *http.Request) (interface{}, int, error) {
		return nil, http.StatusInternalServerError, localerrors.ErrInternal
	}
	handlerFunc := HandlerError(endpointFunc)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), localerrors.ErrInternal.Error())
}

func Test_HandlerError_BadRequest(t *testing.T) {
	assert := assert.New(t)
	endpointFunc := func(_ http.ResponseWriter, _ *http.Request) (interface{}, int, error) {
		return nil, http.StatusBadRequest, errors.New("bad request")
	}
	handlerFunc := HandlerError(endpointFunc)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "bad request")
}

func Test_HandlerError_Pass(t *testing.T) {
	assert := assert.New(t)
	type testBody struct {
		Id int
	}
	expectedObj := testBody{
		Id: 2,
	}
	endpointFunc := func(_ http.ResponseWriter, _ *http.Request) (interface{}, int, error) {
		return expectedObj, http.StatusCreated, nil
	}
	handlerFunc := HandlerError(endpointFunc)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.GreaterOrEqual(res.Code, http.StatusOK)
	assert.LessOrEqual(res.Code, 299)

	receivedObj := testBody{}
	json.Unmarshal(res.Body.Bytes(), &receivedObj)
	assert.Equal(expectedObj, receivedObj)
}
