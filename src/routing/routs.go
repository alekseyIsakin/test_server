package routs

import (
	"net/http"
	"test_server/src/config"
	"test_server/src/model"
	"test_server/src/tokens"

	"github.com/gin-gonic/gin"
)

func UserAuthHandler(c *gin.Context) {
	cfg := config.GetConfig()
	uuid := c.Param("uuid")
	ta, err_a := tokens.GenAccesToken(c, uuid)
	tr, err_r := tokens.GenRefreshToken(c, uuid)

	model.UpdateRefreshTokenForUser(c, tr, uuid)

	if (err_a != nil) || (err_r != nil) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can`t create token"})

		if err_a != nil {
			panic(err_a)
		} else if err_r != nil {
			panic(err_r)
		}
	}

	c.JSON(http.StatusOK, ta+cfg.GetTokenDelimiter()+tr)
}
