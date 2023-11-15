package controllers

import (
	"basic-trade-api/database"
	"basic-trade-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// products
func GetAllProducts(ctx *gin.Context) {
	var products *[]models.Product
	// get all products based on admin ID?
	res := connectDB.DB.Preload("Variant").Find(&products)
	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error occured during data retrieval",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    products,
	})
}

func CreateProduct(ctx *gin.Context) {
	var product *models.Product
	err := ctx.ShouldBindJSON(product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	res := connectDB.DB.Create(&product)
	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create record",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    product,
	})
}

func GetProductByID(ctx *gin.Context) {
	productUUID := ctx.Param("productUUID")
	var product *models.Product
	res := database.ConnectToDB().First(&product, "uuid = ?", productUUID)
	if res != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": res.Error,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})
}

func UpdateProductByID(ctx *gin.Context) {
	productUUID := ctx.Param("productUUID")
	var productUpdated *models.Product
	var product *models.Product
	if err := ctx.ShouldBindJSON(&productUpdated); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	res := database.ConnectToDB().Find(&product, "uuid = ?", productUUID)
	if res.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "No records with given uuid found",
		})
		return
	}

	// delete related variants
	var variants *[]models.Variant
	res = database.ConnectToDB().Delete(&variants, "product_id = ?", product.ID)
	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": res.Error,
		})
		return
	}

	product.ImageURL = productUpdated.ImageURL
	product.ProductName = productUpdated.ProductName
	product.Variants = productUpdated.Variants
	res = database.ConnectToDB().Save(&product)
	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": res.Error,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})
}

func DeleteProductByID(ctx *gin.Context) {
	productUUID := ctx.Param("productUUID")
	var deletedProduct *models.Product
	res := database.ConnectToDB().Clauses(clause.Returning{}).Delete(&deletedProduct, "uuid = ?", productUUID)
	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": res.Error.Error(),
		})
		return
	}
	if res.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "Product with given uuid is not found!",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": deletedProduct,
	})
}

// variants
func GetAllVariants(ctx *gin.Context) {
	return
}

func GetVariantByID(ctx *gin.Context) {
	return
}

func CreateVariant(ctx *gin.Context) {
	return
}

func UpdateVariantByID(ctx *gin.Context) {
	return
}

func DeleteVariantByID(ctx *gin.Context) {
	return
}
