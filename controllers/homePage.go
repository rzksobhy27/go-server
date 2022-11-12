package controllers

import (
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title": "index",
	})
}
