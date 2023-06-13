package main

import (
	"github.com/gin-gonic/gin"
	"go-jwt/initializers"
	"net/http"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	g := gin.Default()

	g.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})

	g.Run(os.Getenv("PORT"))
}
