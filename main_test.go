package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMessageHandler(t *testing.T) {
	router := setupRouter()

	t.Run("Success", func(t *testing.T) {
		input := InputMessage{Message: "Hello world"}
		jsonBody, _ := json.Marshal(input)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/message", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response InputMessage
		_ = json.Unmarshal(w.Body.Bytes(), &response)

		expected := InputMessage{Message: "dlrow olleH"}
		assert.Equal(t, expected, response)
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/message", bytes.NewBuffer([]byte("{invalid json}")))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/message", func(c *gin.Context) {
		var input InputMessage
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		reversedMessage := reverseString(input.Message)
		c.JSON(200, InputMessage{Message: reversedMessage})
	})
	return router
}
