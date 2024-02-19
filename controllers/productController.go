package controllers

import (
	"basic-trade-api/database"
	"basic-trade-api/helpers"
	"basic-trade-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm/clause"
)

// products
func GetAllProducts(ctx *gin.Context) {
	var products []models.Product
	productName := ctx.Query("product_name")
	limit := ctx.Query("limit")
	lastID := ctx.Query("last_prev_id")

	db := database.ConnectToDB()
	query := db.Model(&models.Product{}).Preload("Variants")

	if productName != "" {
		query = query.Where("product_name ILIKE ?", "%"+productName+"%")
	}
	if limit != "" && lastID != "" {
		limitINT, _ := strconv.Atoi(limit)
		query = query.Where("id > ?", lastID).Order("id ASC").Limit(limitINT)
	}

	if err := query.Find(&products).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
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
	// TODO: tambah validasi untuk membatasi ekstensi file image saja yang dapat diupload
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
		adminData := ctx.MustGet("AdminData").(jwt.MapClaims)
		adminID := adminData["id"].(float64)
		helpers.CheckError(err)
		fileName := helpers.RemoveExtension(productForm.ImageFile.Filename)

		// validate file extension
		isImage := helpers.IsImageFile(productForm.ImageFile.Filename)
		if !isImage {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": "Invalid image file extension",
			})
			return
		}

		secureURL, err := helpers.UploadFile(productForm.ImageFile, fileName)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "failed",
				"message": err.Error(),
			})
			return
		}
		product.ProductName = productForm.ProductName
		// product.AdminID = productForm.AdminID
		product.AdminID = int(adminID)
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
	var secureURL string
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

		// validate file extension
		isImage := helpers.IsImageFile(productUpdated.ImageFile.Filename)
		if !isImage {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": "Invalid image file extension",
			})
			return
		}

		secureURL, err = helpers.UploadFile(productUpdated.ImageFile, fileName)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "failed",
				"message": err.Error(),
			})
			return
		}
		// product.ProductName = productUpdated.ProductName
		// product.AdminID = productUpdated.AdminID
		// product.ImageURL = secureURL
		// // check if Variants field is filled in json request or not
		// // if empty, skip variants field assigntment
		// if len(productUpdated.Variants) != 0 {
		// 	product.Variants = productUpdated.Variants
		// }
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
	product.ImageURL = secureURL
	// check if Variants field is filled in json request or not
	// if empty, skip variants field assigntment
	if len(productUpdated.Variants) != 0 {
		product.Variants = productUpdated.Variants
	}

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
