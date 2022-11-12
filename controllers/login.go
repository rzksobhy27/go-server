package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rzksobhy27/go-server/inits"
	"github.com/rzksobhy27/go-server/tables"
	"gorm.io/gorm"
)

func UserLogin(c *gin.Context) {
	body := tables.User{}
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		c.Status(200)
		return
	}

	var user tables.User
	result := inits.DB.Where("name = ?", body.Name).Select("password").First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(200, gin.H{
			"err": "username not found",
		})
		return
	}

	if body.Password != user.Password {
		c.JSON(200, gin.H{
			"err": "passwords doesn't match!",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":  "logged in successfully",
		"user": body,
	})
}
