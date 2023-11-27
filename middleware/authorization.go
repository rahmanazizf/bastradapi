package middleware

import (
	"basic-trade-api/database"
	"basic-trade-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

// function to access product
func AuthorizeProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product *models.Product
		// claim from authentication process
		adminData := ctx.MustGet("AdminData").(jwt5.MapClaims)
		adminID := uint(adminData["id"].(float64))
		// access database
		res := database.ConnectToDB().Find(&product, "uuid = ?", ctx.Param("productUUID"))
		if res.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "failed",
				"message": res.Error.Error(),
			})
		}
		if res.RowsAffected == 0 {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "failed",
				"message": "No records with given uuid found",
			})
			return
		}
		// query data
		if product.AdminID != int(adminID) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "You are unauthorized",
			})
			return
		}
		// if the data is matched with claims, return the data
		ctx.Next()
	}
}

func AuthorizeVariant() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var variant *models.Variant
		var product *models.Product
		// claim from authentication process
		adminData := ctx.MustGet("AdminData").(jwt5.MapClaims)
		adminID := uint(adminData["id"].(float64))
		// access database
		res := database.ConnectToDB().Find(&variant, "uuid = ?", ctx.Param("variantUUID"))
		if res.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "failed",
				"message": res.Error.Error(),
			})
		}
		if res.RowsAffected == 0 {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "failed",
				"message": "No records with given uuid found",
			})
			return
		}
		// query data
		database.ConnectToDB().Find(&product, "id = ?", variant.ProductID)
		if product.AdminID != int(adminID) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "You are unauthorized",
			})
			return
		}
		// if the data is matched with claims, return the data
		ctx.Next()
	}
}
