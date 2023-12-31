package controllers

import (
	"basic-trade-api/database"
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
				"status":  "failed",
				"message": err.Error(),
			})
			return
		}

	} else {
		err := ctx.ShouldBind(&admin)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": err.Error(),
			})
			return
		}
	}
	// check if the same email had already been registered
	res := database.ConnectToDB().First(&admin, "email = ?", admin.Email)
	if res.RowsAffected > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "The same email had already been registered. Please use a different email.",
		})
		return
	}
	// res = connectDB.DB.Create(&admin)
	res = database.ConnectToDB().Create(&admin)
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "No rows created in the database",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   admin,
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
				"status":  "failed",
				"message": err.Error(),
			})
			return
		}

	} else {
		err := ctx.ShouldBind(&adminInput)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": err.Error(),
			})
			return
		}
	}
	res := database.ConnectToDB().First(&admin, "email = ?", adminInput.Email)
	compareUsername := admin.Name == adminInput.Name
	isPwdCorrect := helpers.ComparePwd(admin.Password, adminInput.Password, admin.Salt)
	if res.RowsAffected > 0 && compareUsername && isPwdCorrect {
		token := helpers.GenerateToken(admin.ID, admin.Email)
		ctx.JSON(http.StatusOK, gin.H{
			"status": "Login successful",
			"token":  token,
		})
		return
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "Invalid username or password",
		})
	}
}
