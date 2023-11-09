package routers

import (
	"basic-trade-api/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello, Mom!",
		})
	})
	// admin authentication
	auth := router.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}
	// products
	products := router.Group("/products")
	{
		// products query
		products.POST("/")
		products.GET("/")
		products.GET("/:productUUID")
		products.PUT("/:productUUID")
		products.DELETE("/:productUUID")
		// variants query
		products.GET("/variants")
		products.GET("/variants/:variantUUID")
		products.POST("/variants")
		products.PUT("/variants/:variantUUID")
		products.DELETE("/variants/:variantUUID")
	}

	return router
}
