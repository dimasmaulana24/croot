package url

import (
	"github.com/InformaticsResearchCenter/croot/controller"
	"github.com/gin-gonic/gin"
)

func Site(page *gin.Engine) {
	page.GET("/", controller.Get)
	page.POST("/api", controller.PostApi)
	page.POST("/m/gis/SG/airport", controller.PostGisSG)
}
