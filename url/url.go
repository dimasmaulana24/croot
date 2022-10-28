package url

import (
	"github.com/InformaticsResearchCenter/croot/controller"
	"github.com/gin-gonic/gin"
)

func Page(site *gin.Engine) {
	site.GET("/", controller.Get)
	site.POST("/api", controller.PostApi)
}
