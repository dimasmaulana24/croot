package controller

import "github.com/gin-gonic/gin"

func Get(c *gin.Context) {
	html := "selamat datang <b>Abang Ganteng</b>"
	c.JSON(200, html)
}
