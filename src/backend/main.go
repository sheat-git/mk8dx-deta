package main

import (
	"encoding/json"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sheat-git/mkdeta/controller"
)

func main() {
	r := gin.Default()
	var oris = []string{}
	err := json.Unmarshal([]byte(os.Getenv("CORS_ORIGINS")), &oris)
	if err != nil {
		panic(err)
	}
	r.Use(cors.New(cors.Config{
		AllowOrigins: oris,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))
	r.GET("/sokuji", controller.SokujiGet)
	r.Run()
}
