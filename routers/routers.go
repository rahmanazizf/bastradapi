package routers

import (
	"basic-trade-api/controllers"
	"basic-trade-api/middleware"
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
		products.GET("/", controllers.GetAllProducts)
		products.GET("/:productUUID", controllers.GetProductByID)
		// variants query
		products.GET("/variants")
		products.GET("/variants/:variantUUID")
		products.Use(middleware.Authenticate())
		products.POST("/", controllers.CreateProduct)
		products.PUT("/:productUUID", middleware.AuthorizedProduct(), controllers.UpdateProductByID)
		products.DELETE("/:productUUID", middleware.AuthorizedProduct(), controllers.DeleteProductByID)
		products.POST("/variants")
		products.PUT("/variants/:variantUUID")
		products.DELETE("/variants/:variantUUID")
	}

	return router
}
