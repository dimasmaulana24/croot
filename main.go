package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/informatics-research-center/croot/url"
)

func main() {
	port := os.Getenv("PORT")
	SetPort(&port)

	web := gin.New()
	web.SetTrustedProxies([]string{"127.0.0.1"})

	url.Page(web)
	web.Run(":" + port)
}

func SetPort(port *string) {
	if len(*port) == 0 {
		*port = "80"
	}
}
