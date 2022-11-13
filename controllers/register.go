package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rzksobhy27/go-server/inits"
	"github.com/rzksobhy27/go-server/tables"
	"gorm.io/gorm"
)

func UserRegister(c *gin.Context) {
	body := tables.User{}
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		c.Status(http.StatusOK)
		return
	}

	result := inits.DB.Where("name = ?", body.Name).Select("").First(&tables.User{})

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"err": "username already exists!",
		})
		return
	}

	inits.DB.Create(body)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":     body.Name,
		"password": body.Password,
	})

	key := os.Getenv("TOKEN_KEY")
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		c.Status(http.StatusOK)
		return
	}

	c.SetCookie("token", tokenString, 30*60*1000, "/", "localhost", false, false)

	c.Status(http.StatusOK)
}
