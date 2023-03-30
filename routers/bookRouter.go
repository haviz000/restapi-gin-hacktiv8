package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/haviz000/restapi-gin-hacktiv8/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controllers.GetBooks)
	router.GET("/book/:idBook", controllers.GetBookById)
	router.POST("/book", controllers.CreateBook)
	router.PUT("/book/:idBook", controllers.UpdateBook)
	router.DELETE("/book/:idBook", controllers.DeleteBook)

	return router
}
