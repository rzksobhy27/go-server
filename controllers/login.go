package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/rzksobhy27/go-server/inits"
	"github.com/rzksobhy27/go-server/tables"
	"gorm.io/gorm"
)

func UserLogin(c *gin.Context) {
	body := tables.User{}
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		c.Status(http.StatusOK)
		return
	}

	var user tables.User
	result := inits.DB.Where("name = ?", body.Name).Select("password").First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"err": "username not found",
		})
		return
	}

	if body.Password != user.Password {
		c.JSON(http.StatusOK, gin.H{
			"err": "passwords doesn't match!",
		})
		return
	}

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
