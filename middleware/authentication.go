package middleware

import (
	"basic-trade-api/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// authenticate user and returns gin handler function that will be passed to gin routing
func Authenticate() gin.HandlerFunc {
	// return gin handler func to be used by routes
	return func(ctx *gin.Context) {
		res, err := helpers.VerifyToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "Unauthenticated",
				"message": err.Error(),
			})
			return
		}
		// set claims to context with key "UserData"
		ctx.Set("UserData", res)
		// proceed to handler function
		ctx.Next()

	}
}
