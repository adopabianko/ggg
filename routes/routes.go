package routes

import (
	"github.com/adopabianko/ggg-boilerplate/controllers"
	"github.com/gin-gonic/gin"
)

// Setup Router
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/book-store")
	{
		grp1.GET("book", controllers.GetBook)
		grp1.POST("book", controllers.CreateBook)
		grp1.GET("book/:id", controllers.GetBookByID)
		grp1.PUT("book/:id", controllers.UpdateBook)
		grp1.DELETE("book/:id", controllers.DeleteBook)
	}

	return r
}
