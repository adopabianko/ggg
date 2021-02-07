package controllers

import (
	"net/http"

	"github.com/adopabianko/ggg-boilerplate/models"

	"github.com/gin-gonic/gin"
)

// Get All Books
func GetBook(c *gin.Context) {
	var book []models.Book

	err := models.GetAllBooks(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
		})
	} else {
		if len(book) > 0 {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "List of books",
				"data":    book,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "Empty data",
			})
		}

	}
}

// Create Book
func CreateBook(c *gin.Context) {
	var book models.Book

	c.BindJSON(&book)
	err := models.CreateBook(&book)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "The book successfully saved",
			"data":    book,
		})
	}
}

// Get the book by id
func GetBookByID(c *gin.Context) {
	id := c.Params.ByName("id")

	var book models.Book

	err := models.GetBookByID(&book, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "The book data",
			"data":    book,
		})
	}
}

// Update the book information
func UpdateBook(c *gin.Context) {
	var book models.Book

	id := c.Params.ByName("id")

	err := models.GetBookByID(&book, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Not found",
		})
	} else {
		c.BindJSON(&book)

		err = models.UpdateBook(&book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "The book successfully updated",
				"data":    book,
			})
		}
	}
}

// Delete the book
func DeleteBook(c *gin.Context) {
	var book models.Book

	id := c.Params.ByName("id")

	err := models.GetBookByID(&book, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Not found",
		})
	} else {
		err = models.DeleteBook(&book, id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "The book successfully deleted",
			})
		}
	}
}
