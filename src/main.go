package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/junxxx/notion-book-creater/service"
)

type book struct {
	Isbn string `json:"isbn"`
}

const apiKeyEnv = "BOOK_CREATOR_API_KEY"

func postBook(c *gin.Context) {
	var b book
	if err := c.BindJSON(&b); err != nil {
		log.Println(err)
		return
	}
	err := service.CreatePage(b.Isbn)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusCreated, b)
}

// api auth
func validAPIKey(c *gin.Context) {
	apiKey := c.Request.Header.Get("x-api-key")
	if len(apiKey) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}
	if apiKey != os.Getenv(apiKeyEnv) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}
	c.Next()
}

func main() {
	router := gin.New()
	b := router.Group("/book")
	b.Use(validAPIKey)
	{
		b.POST("/create", postBook)
	}
	router.Run(":8080")
}
