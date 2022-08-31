package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ashishbhatt01/registeryApp/app/models"
	testhelper "github.com/ashishbhatt01/registeryApp/app/test-helper"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	testhelper.IntializeTestLogger()
	service := testhelper.MockRegisterService{}
	ctrl := NewRegistryController(service)

	ginEngine, req, respRecorder := createHandleSetup(t, "POST", "/registery/add", ctrl.Add, bytes.NewBuffer([]byte(`{"Value":10}`)))
	ginEngine.ServeHTTP(respRecorder, req)
	assert.Equal(t, http.StatusOK, respRecorder.Code, "it should return status OK")
}

func Test_Subs(t *testing.T) {
	testhelper.IntializeTestLogger()
	service := testhelper.MockRegisterService{}
	ctrl := NewRegistryController(service)

	ginEngine, req, respRecorder := createHandleSetup(t, "POST", "/registery/subs", ctrl.Subs, bytes.NewBuffer([]byte(`{"Value":10}`)))
	ginEngine.ServeHTTP(respRecorder, req)
	assert.Equal(t, http.StatusOK, respRecorder.Code, "it should return status OK")
}

func Test_Get(t *testing.T) {
	sampleData := models.Register{
		Value: 100,
	}
	service := testhelper.MockRegisterService{}
	service.SetData(sampleData)
	ctrl := NewRegistryController(service)

	ginEngine, req, respRecorder := createHandleSetup(t, "GET", "/registery/value", ctrl.Get, nil)
	ginEngine.ServeHTTP(respRecorder, req)

	var data models.Register
	dataString := respRecorder.Body.String()
	json.Unmarshal([]byte(dataString), &data)

	assert.Equal(t, http.StatusOK, respRecorder.Code, "it should return status OK")
	assert.NotNil(t, data, "data should not be nil")
	assert.Equal(t, 100, data.Value, "received value should be 100")
}

func createHandleSetup(t *testing.T, method string, endPoint string, handler gin.HandlerFunc, body io.Reader) (*gin.Engine, *http.Request, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	ginEngine := gin.Default()
	var httpRequest *http.Request
	var err error

	if method == "POST" {
		ginEngine.POST(endPoint, handler)
		httpRequest, err = http.NewRequest("POST", endPoint, body)
	} else {
		ginEngine.GET(endPoint, handler)
		httpRequest, err = http.NewRequest("GET", endPoint, body)
	}

	if err != nil {
		t.Fatal(err)
	}
	requestRecorder := httptest.NewRecorder()

	return ginEngine, httpRequest, requestRecorder
}
