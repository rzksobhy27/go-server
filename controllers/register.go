package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rzksobhy27/go-server/inits"
	"github.com/rzksobhy27/go-server/tables"
	"gorm.io/gorm"
)

func UserRegister(c *gin.Context) {
	body := tables.User{}
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		c.Status(200)
		return
	}

	result := inits.DB.Where("name = ?", body.Name).Select("").First(&tables.User{})

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(200, gin.H{
			"err": "username already exists!",
		})
		return
	}

	inits.DB.Create(body)

	c.JSON(200, gin.H{
		"msg": "user created successfully",
	})

	c.Status(200)
}
