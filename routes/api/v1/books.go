package v1

import (
	
	"github.com/gin-gonic/gin"

	v1 "golang_rest_api/controllers/api/v1"
)

func InitRouter(){
	router := gin.Default()
	v1Router := router.Group("/v1")
	{
		v1Router.GET("/books",v1.GetBooks)
		v1Router.GET("/books/:id",v1.GetBookByID)
		v1Router.POST("/books",v1.PostBook)
	}
	router.Run("localhost:8080")

}