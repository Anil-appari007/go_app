package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	// "github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestSayHello(t *testing.T) {

	router := SetUpRouter()
	router.GET("/hello", sayHello)
	req, _ := http.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	actualResponseData, _ := io.ReadAll(w.Body)

	// expectedResponse := `{"message:"hello2"}`
	expectedResponse := `{"message":"hello2"}`
	assert.Equal(t, expectedResponse, string(actualResponseData))
	assert.Equal(t, 200, w.Code)

}
