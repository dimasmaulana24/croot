package url

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/informatics-research-center/croot/controller"
)

func Page(site *gin.Engine) {
	site.GET("/", controller.Get)
	site.POST("/api", controller.PostApi)
}
