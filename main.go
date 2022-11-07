package main

import (
	"os"

	"github.com/InformaticsResearchCenter/croot/url"
	"github.com/gin-gonic/gin"
)

func main() {
	web := gin.New()
	web.SetTrustedProxies(nil)

	url.Site(web)
	web.Run(":" + SetPort())
}

func SetPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "80"
	}
	return port
}
