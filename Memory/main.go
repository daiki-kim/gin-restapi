package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Rating float64 `json:"rating"`
}

var books = []book{
	{ID: "1", Title: "Atomic Habits", Author: "James Clear", Rating: 4.9},
	{ID: "2", Title: "The Almanack of Naval Ravikant", Author: "Eric Jorgenson", Rating: 5.0},
	{ID: "3", Title: "The Psychology of Money", Author: "Morgan Housel", Rating: 4.6},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postBooks)
	router.PUT("books/:id", updateBookByID)
	router.DELETE("books/:id", deleteBookByID)

	router.Run("localhost:8000")
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBooks(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func updateBookByID(c *gin.Context) {
	id := c.Param("id")
	var updatedBook book
	if err := c.BindJSON(&updatedBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error updating book"})
		return
	}
	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			c.IndentedJSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func deleteBookByID(c *gin.Context) {
	id := c.Param("id")
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
