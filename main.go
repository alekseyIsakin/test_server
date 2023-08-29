package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"test_server/src/model"
	"test_server/src/routing"
)

type res struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
}

func mainGetHandler(c *gin.Context) {
	c.JSON(200, res{Value: 10, Name: "Hello"})
}

func mainPostHandler(c *gin.Context) {
	r := new(res)
	if err := c.Bind(r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, nil)

	fmt.Printf("\n%s, %d\n", r.Name, r.Value)
}

func main() {
	// model.SetupExampleData(context.TODO())
	router := gin.Default()
	a := routs.Authorizer{}
	router.GET("/main", a.UserAuthHandler)
	router.POST("/main", mainPostHandler)
	router.Run(":8000")
}
