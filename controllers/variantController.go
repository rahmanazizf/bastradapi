package controllers

import (
	"basic-trade-api/database"
	"basic-trade-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// variants
func GetAllVariants(ctx *gin.Context) {
	var variants *[]models.Variant
	res := database.ConnectToDB().Find(&variants)
	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": res.Error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   variants,
	})
}

func GetVariantByID(ctx *gin.Context) {
	variantUUID := ctx.Param("variantUUID")
	var variant *models.Variant
	res := database.ConnectToDB().Find(&variant, "uuid = ?", variantUUID)
	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": res.Error.Error(),
		})
		return
	}
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "not found",
			"message": "no records with given uuid is found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   variant,
	})
}

func CreateVariant(ctx *gin.Context) {
	var variant *models.Variant
	err := ctx.ShouldBindJSON(&variant)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}
	res := database.ConnectToDB().Create(&variant)
	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": res.Error.Error(),
		})
		return
	}
	if res.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNoContent, gin.H{
			"status":  "no content",
			"message": "no records are created",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   variant,
	})
}

func UpdateVariantByID(ctx *gin.Context) {
	variantUUID := ctx.Param("variantUUID")
	var variantUpdated *models.Variant
	var variant *models.Variant
	err := ctx.ShouldBindJSON(&variantUpdated)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}
	res := database.ConnectToDB().First(&variant, "uuid = ?", variantUUID)
	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": res.Error.Error(),
		})
		return
	}
	if res.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "not found",
			"message": "no records with given uuid found",
		})
		return
	}

	return
}

func DeleteVariantByID(ctx *gin.Context) {
	return
}
