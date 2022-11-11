package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rzksobhy27/go-server/inits"
	"github.com/rzksobhy27/go-server/tables"
)

func UserRegister(c *gin.Context) {
	user := tables.User{Name: "rzksobhy27", Password: "12345678"}
	result := inits.DB.Create(&user)

	if result.Error != nil {
		c.Status(200)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}
