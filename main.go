package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "The Last Born", Author: "Jen Doe", Quantity: 5},
	{ID: "2", Title: "The Deffender", Author: "Planaf Kajal", Quantity: 5},
	{ID: "3", Title: "Money Buy Happiness", Author: "Jen Doe", Quantity: 5},
	{ID: "4", Title: "Richest Man", Author: "Zubery Pima", Quantity: 5},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {

	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/cbooks", createBook)
	router.Run("localhost:8081")

}
