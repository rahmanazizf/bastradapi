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
	// parse token
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("You are not logged in")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, errors.New("You are not logged in")
	}
	// convert parsedToken to jwt claim
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok && !parsedToken.Valid {
		return nil, errors.New("You are not logged in")
	}
	expClaim, exists := claims["expired"]
	if !exists {
		return nil, errors.New("Expire claim is missing")
	}
	// ensure expClaim is string
	expStr, ok := expClaim.(string)
	if !ok {
		return nil, errors.New("Expire claim is not a valid type")
	}
	// parse time
	expTime, err := time.Parse(time.RFC3339, expStr)
	if err != nil {
		return nil, errors.New("Error parsing expiration time")
	}
	// ensure current time is not beyond expire date
	if time.Now().After(expTime) {
		return nil, errors.New("Token is expired")
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}
