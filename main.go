package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

	fmt.Printf("\n%s, %d, %s\n", r.Name, r.Value)
}

func main() {
	router := gin.Default()
	router.GET("/main", mainGetHandler)
	router.POST("/main", mainPostHandler)
	router.Run(":8000")
}
