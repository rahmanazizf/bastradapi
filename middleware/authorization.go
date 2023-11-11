package middleware

import "github.com/gin-gonic/gin"

// function to access product
func AuthorizedProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// access params

		// access database

		// get claims from authentication process

		// query data

		// match claims with the data had been queried

		// if the data is matched with claims, return the data

		// else, abort
	}
}
