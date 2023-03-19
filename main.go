package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type InputMessage struct {
	Message string `json:"message"`
}

func main() {
	router := gin.Default()

	router.POST("/message", func(c *gin.Context) {
		var input InputMessage
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		log.Printf("Info: Received message: %s", input.Message)

		reversedMessage := reverseString(input.Message)
		c.JSON(200, InputMessage{Message: reversedMessage})
	})

	router.Run(":8080")
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
