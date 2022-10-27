package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/informatics-research-center/croot/url"
)

func main() {
	port := os.Getenv("PORT")
	web := gin.New()
	url.Page(web)
	web.Run(":" + port)
}
