package main

import (
	"gitlab.com/informatics-research-center/croot/url"

	"github.com/gin-gonic/gin"
)

func main() {
	web := gin.New()
	url.Page(web)
	web.Run(":5000")
}
