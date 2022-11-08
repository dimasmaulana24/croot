package controller

import (
	"github.com/InformaticsResearchCenter/croot/helper"
	"github.com/InformaticsResearchCenter/croot/model"
	_ "github.com/InformaticsResearchCenter/croot/module/gis/SG"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var response model.Response
	response.Response = "selamat datang Abang Ganteng IP Server " + helper.GetIPaddress()
	c.JSON(200, &response)
}

func PostApi(c *gin.Context) {
	var chat model.Chat
	c.BindJSON(&chat)
	var response model.Response
	response.Response = "selamat datang " + chat.Phone_number + " " + chat.Messages
	c.JSON(200, &response)
}

func PostGisSG(c *gin.Context) {
	var gis model.Gis
	c.BindJSON(&gis)
	helper.Call("sg.SetAirport", &gis)
	//sg.SetAirport(&gis)
	var response model.Response
	response.Response = "shp file " + gis.Name + " was created "
	c.JSON(200, &response)
}
