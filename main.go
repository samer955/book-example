package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []book{

	{ID: "1", Title: "Science of Logic", Author: "Georg Wilhelm Friedrich Hegel"},
	{ID: "2", Title: "The Wisdom of Life", Author: "Arthur Schopenhauer"},
	{ID: "3", Title: "The Metamorphosis", Author: "Franz Kafka"},
	{ID: "4", Title: "Das Kapital", Author: "Karl Marx"},
	{ID: "5", Title: "The Sorrows of Young Werther", Author: "Johann Wolfgang von Goethe"},
}

func main() {

	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postBooks)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8182"
	}

	httpServer := os.Getenv("HTTP_SERVER")
	if httpServer == "" {
		httpServer = "localhost"
	}

	router.Run(httpServer + ":" + httpPort)

}

func getBooks(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, books)

}

func postBooks(c *gin.Context) {

	var newBook book

	// Call BindJSON to bind the received JSON to newBook.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

func getBookByID(c *gin.Context) {

	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})

}
