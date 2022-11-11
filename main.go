package main

import (
	"github.com/rzksobhy27/go-server/controllers"
	"github.com/rzksobhy27/go-server/inits"

	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()
	inits.ConnectDB()
}

func main() {
	r := gin.Default()

	r.POST("/", controllers.UserRegister)

	r.Run()
}
