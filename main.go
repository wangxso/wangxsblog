package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wangxso/wangxsoblog/routes"
	"github.com/wangxso/wangxsoblog/utils"
)

func main() {
	r := gin.Default()
	config, err := utils.GetConfig()
	if err != nil {
		log.Fatal("error at read config, err, ", err)
	}
	routes.SetupApiRoutes(r)
	routes.SetupAuthRoutes(r)

	r.Run(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port))
}
