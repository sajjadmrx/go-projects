package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-jwt/initializers"
	"go-jwt/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		ResponseErrorWrapper(c, http.StatusBadRequest, "Failed to read body")
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		ResponseErrorWrapper(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		ResponseErrorWrapper(c, http.StatusInternalServerError, "Failed to save user")
		return
	}

	// create Jwt
	jToken, err := createJwt(user.ID)
	if err != nil {
		ResponseErrorWrapper(c, http.StatusInternalServerError, "Failed to create Token")
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": jToken,
	})
}

func ResponseErrorWrapper(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"message": message,
	})
}

func createJwt(userId uint) (string, error) {
	mySigningKey := []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString(mySigningKey)
	return tokenString, err
}