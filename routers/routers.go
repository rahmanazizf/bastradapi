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
		products.GET("/variants", controllers.GetAllVariants)
		products.GET("/variants/:variantUUID", controllers.GetVariantByID)
		products.Use(middleware.Authenticate())
		products.POST("/", controllers.CreateProduct)
		products.PUT("/:productUUID", middleware.AuthorizeProduct(), controllers.UpdateProductByID)
		products.DELETE("/:productUUID", middleware.AuthorizeProduct(), controllers.DeleteProductByID)
		products.POST("/variants", controllers.CreateVariant)
		products.PUT("/variants/:variantUUID", middleware.AuthorizeVariant(), controllers.UpdateVariantByID)
		products.DELETE("/variants/:variantUUID", middleware.AuthorizeVariant(), controllers.DeleteVariantByID)
	}

	return router
}
