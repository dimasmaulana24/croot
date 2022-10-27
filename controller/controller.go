package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/informatics-research-center/croot/model"
)

func Get(c *gin.Context) {
	html := "selamat datang <b>Abang Ganteng</b>"
	c.JSON(200, html)
}

func PostApi(c *gin.Context) {
	var chat model.Chat
	c.BindJSON(&chat)
	var response model.Response
	response.Response = "selamat datang " + chat.Phone_number + " " + chat.Messages
	c.JSON(200, &response)
}
