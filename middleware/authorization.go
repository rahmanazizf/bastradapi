package middleware

import (
	"basic-trade-api/database"
	"basic-trade-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// function to access product
func AuthorizedProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product *models.Product
		// access params
		userData, exists := ctx.Get("userData")
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not an authorized user!",
				"data":    userData, // TO DELETE
			})
		}
		// access database
		database.ConnectToDB().Find(&product)
		// get claims from authentication process

		// query data

		// match claims with the data had been queried

		// if the data is matched with claims, return the data

		// else, abort
	}
}
