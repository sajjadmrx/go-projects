package main

import (
	"github.com/gin-gonic/gin"
	"go-jwt/controllers"
	"go-jwt/initializers"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	g := gin.Default()
	//auth := g.Group("/auth")

	g.POST("/signup", controllers.Signup)

	g.Run(os.Getenv("PORT"))

}
