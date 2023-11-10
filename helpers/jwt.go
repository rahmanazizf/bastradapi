package helpers

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// generate token
// generate token dengan menggunakan credential dan secret key
var secretKey = "this-is-mY-s3c12eT-k3Y"

func GenerateToken(userID int, email string) string {
	// buat klaim
	claims := jwt.MapClaims{
		"id":      userID,
		"email":   email,
		"expired": time.Now().Add(time.Minute * 1),
	}

	// buat token dengan menggunakan klaim
	parsedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// tanda tangan token yang sudah dibuat
	signedToken, err := parsedToken.SignedString([]byte(secretKey))
	CheckError(err)
	// return
	return signedToken
}

// verify token
func VerifyToken(ctx *gin.Context) (interface{}, error) {
	// get header request and extract from authorization key
	auth := ctx.Request.Header.Get("Authorization")
	// get bearer token
	bearerExist := strings.HasPrefix(auth, "Bearer")
	if !bearerExist {
		return nil, errors.New("You are not logged in")
	}
	token := strings.Split(auth, "")[1]
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, errors.New("You are not logged in")
	}
	return parsedToken, nil
}
