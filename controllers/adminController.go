package controllers

import (
	"basic-trade-api/helpers"
	"basic-trade-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var contentType = "application/json"

func Register(ctx *gin.Context) {
	var admin models.Admin
	if contentType == ctx.ContentType() {
		err := ctx.ShouldBindJSON(&admin)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
				"error":   err.Error(),
			})
			return
		}

	} else {
		err := ctx.ShouldBind(&admin)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
				"error":   err.Error(),
			})
			return
		}
	}
	res := connectDB.DB.Create(&admin)
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "No rows created in the database",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    admin,
		"message": "success",
	})

}

func Login(ctx *gin.Context) {
	// TODO: login mengimplementasi searching data user di database dengan
	// token untuk authorization juga diperoleh dari proses autentikasi/login
	var adminInput *models.Admin
	var admin *models.Admin
	if contentType == ctx.ContentType() {
		err := ctx.ShouldBindJSON(&adminInput)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
				"error":   err.Error(),
			})
			return
		}

	} else {
		err := ctx.ShouldBind(&adminInput)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
				"error":   err.Error(),
			})
			return
		}
	}
	res := connectDB.DB.First(&admin)
	compareUsername := admin.Name == adminInput.Name
	isPwdCorrect := helpers.ComparePwd(admin.Password, adminInput.Password, admin.Salt)
	if res.RowsAffected > 0 && compareUsername && isPwdCorrect {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"data":    admin,
		})
		return
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Invalid username or password",
		})
	}
}
