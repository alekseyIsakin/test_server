package main

import (
	"context"

	"github.com/gin-gonic/gin"

	"test_server/src/config"
	"test_server/src/model"
	"test_server/src/routing"
)

func main() {
	config.Init()
	model.SetupExampleData(context.TODO())

	router := gin.Default()

	router.GET("/auth/:uuid", routs.UserAuthHandler)
	router.GET("/refr", routs.RenewRefreshToken)

	err := router.Run(":8000")
	if err != nil {
		panic("Cant start server")
	}
}
