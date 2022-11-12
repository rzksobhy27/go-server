package main

import (
	"github.com/gin-gonic/gin"

	"github.com/rzksobhy27/go-server/controllers"
	"github.com/rzksobhy27/go-server/inits"
)

func init() {
	inits.LoadEnv()
	inits.ConnectDB()
}

func main() {
	r := gin.Default()

	api := r.Group("/api")

	api.POST("/login", controllers.UserLogin)
	api.POST("/register", controllers.UserRegister)

	r.Run(":6969")
}
