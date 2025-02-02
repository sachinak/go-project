package main

import (
	"bytes"
	"go-project/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestScanFourFiles(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/scan", api.Scan)

	requestBody := `{
		"repo":"https://raw.githubusercontent.com/velancio/vulnerability_scans/main/",
		"files":["vulnscan18.json", "vulnscan456.json", "vulnscan123.json", "vulnscan789.json"]
	  }`
	req, _ := http.NewRequest(http.MethodPost, "/scan", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status Codes should be same")
}

func TestScanMalformedTwoFiles(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/scan", api.Scan)

	requestBody := `{
		"repo":"https://raw.githubusercontent.com/velancio/vulnerability_scans/main/",
		"files":["vulnscan18.json", "vulnscan789.json"]
	  }`
	req, _ := http.NewRequest(http.MethodPost, "/scan", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status Codes should be same")
}

func TestScanMalformedMalformedGithubUrl(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/scan", api.Scan)

	requestBody := `{
		"repo":"https://raw.gthubusercontent.com/velancio/vulnerability_scans/main/",
		"files":["vulnscan18.json", "vulnscan789.json"]
	  }`
	req, _ := http.NewRequest(http.MethodPost, "/scan", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status Codes should be same")
}

func TestScanMalformedMalformedReqBody(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/scan", api.Scan)

	requestBody := `{
		"repo":"https://raw.githubusercontent.com/velancio/vulnerability_scans/main/",
		"files":["vulnscan18.json", "vulnscan789.json"
	  }`
	req, _ := http.NewRequest(http.MethodPost, "/scan", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Status Codes should be same")
}

func TestQuery(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/query", api.Query)

	requestBody := `{
		"filters":{
			"severity":"HIGH"
		}
	  }`
	req, _ := http.NewRequest(http.MethodPost, "/query", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status Codes should be same")
}

func TestQueryMalformedReqBody(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/query", api.Query)

	requestBody := `{
		"filters":{
			"severity":"HIGH"
		
	  }`
	req, _ := http.NewRequest(http.MethodPost, "/query", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Status Codes should be same")
}
