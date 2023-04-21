package utils

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUpRecorder() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}

func TestGenerateSuccessResp(t *testing.T) {
	recorder := setUpRecorder()
	data := "Test Success"
	successStatusCode := 201

	GenerateJsonResponse(recorder, data, successStatusCode, "Test Success")

	var resp Response
	json.NewDecoder(recorder.Body).Decode(&resp)

	assert.Equal(t, "Test Success", resp.Message)
	assert.Equal(t, HttpMessage[successStatusCode], resp.Status)
	assert.Equal(t, successStatusCode, recorder.Result().StatusCode)
	assert.Equal(t, "Test Success", resp.Data)
}

func TestGenerateErrorResp(t *testing.T) {
	recorder := setUpRecorder()
	data := "Test Failed"
	errorStatusCode := 400

	GenerateJsonResponse(recorder, data, errorStatusCode, "Test Failed")

	var resp Response
	json.NewDecoder(recorder.Body).Decode(&resp)

	assert.Equal(t, "Test Failed", resp.Message)
	assert.Equal(t, HttpMessage[errorStatusCode], resp.Status)
	assert.Equal(t, 400, recorder.Result().StatusCode)
	assert.Equal(t, "Test Failed", resp.Data)
}