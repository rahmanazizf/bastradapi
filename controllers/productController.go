package controllers

import (
	"basic-trade-api/database"
	"basic-trade-api/helpers"
	"basic-trade-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// products
func GetAllProducts(ctx *gin.Context) {
	var products []models.Product
	productName := ctx.Query("product_name")

	db := database.ConnectToDB()
	var res *gorm.DB
	if productName != "" {
		res = db.Preload("Variants").Where("product_name ILIKE ?", "%"+productName+"%").Find(&products)
	} else {
		res = db.Preload("Variants").Find(&products)
	}
	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": res.Error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   products,
	})
}

func CreateProduct(ctx *gin.Context) {
	// var product *models.Product
	// var productForm *models.ProductImage
	product := models.Product{}
	productForm := models.ProductImage{}
	var err error
	if ctx.Request.Header.Get("Content-Type") == "application/json" {
		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": err.Error(),
			})
			return
		}
	} else {
		err = ctx.ShouldBind(&productForm)
		helpers.CheckError(err)
		fileName := helpers.RemoveExtension(productForm.ImageFile.Filename)
		secureURL, err := helpers.UploadFile(productForm.ImageFile, fileName)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "failed",
				"message": err.Error(),
			})
			return
		}
		product.ProductName = productForm.ProductName
		product.AdminID = productForm.AdminID
		product.ImageURL = secureURL
		product.Variants = productForm.Variants
	}

	res := database.ConnectToDB().Create(&product)
	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": res.Error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})
}

func GetProductByID(ctx *gin.Context) {
	productUUID := ctx.Param("productUUID")
	var product *models.Product
	res := database.ConnectToDB().Preload("Variants").First(&product, "uuid = ?", productUUID)
	if res.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "no records with given uuid found",
		})
	}
	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": res.Error.Error(),
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
	var productUpdated = models.ProductImage{}
	var product = models.Product{}
	var err error
	if ctx.Request.Header.Get("Content-Type") == "application/json" {
		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": err.Error(),
			})
			return
		}
	} else {
		err = ctx.ShouldBind(&productUpdated)
		helpers.CheckError(err)
		fileName := helpers.RemoveExtension(productUpdated.ImageFile.Filename)
		secureURL, err := helpers.UploadFile(productUpdated.ImageFile, fileName)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "failed",
				"message": err.Error(),
			})
			return
		}
		product.ProductName = productUpdated.ProductName
		product.AdminID = productUpdated.AdminID
		product.ImageURL = secureURL
		product.Variants = productUpdated.Variants
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
			"message": res.Error.Error(),
		})
		return
	}

	product.ProductName = productUpdated.ProductName
	product.Variants = productUpdated.Variants
	res = database.ConnectToDB().Save(&product)
	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": res.Error.Error(),
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
