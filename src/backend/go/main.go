package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sheat-git/mkdeta/controller"
)

func main() {
	r := gin.Default()
	r.GET("/sokuji", controller.SokujiGet)
	r.Run()
}
