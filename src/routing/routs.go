package routs

import (
	"encoding/base64"
	"net/http"
	"test_server/src/config"
	"test_server/src/model"
	"test_server/src/tokens"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserAuthHandler(c *gin.Context) {
	cfg := config.GetConfig()
	uuid := c.Param("uuid")
	ta, err_a := tokens.GenAccesToken(c, uuid)
	tr, err_r := tokens.GenRefreshToken(c, uuid)
	tr = base64.RawURLEncoding.EncodeToString([]byte(tr))

	model.UpdateRefreshTokenForUser(c, tr, uuid)

	if err_a != nil || err_r != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can`t create token"})
		panic(err_a)
	}

	c.SetCookie(
		"refresh",
		tr,
		cfg.GetMaxAgeRefresh(),
		"/",
		cfg.GetDomain(),
		true,
		true)

	c.JSON(http.StatusOK, ta+cfg.GetTokenDelimiter()+tr)
}

func RenewRefreshToken(c *gin.Context) {
	cfg := config.GetConfig()
	cr, err := c.Cookie("refresh")

	if err != nil {
		c.JSON(http.StatusUnauthorized, "wrong refresh token")
		panic(err)
	}
	token, _ := jwt.DecodeSegment(cr)
	uuid, err := tokens.ValidRefreshToken(c, string(token), []byte(cfg.GetRefreshSecret()))

	if err != nil {
		c.JSON(http.StatusUnauthorized, "cant provide new token")
		panic(err)
	}

	model.ReplaceRefreshTokenForUser(c, string(token), uuid)

	c.JSON(http.StatusOK, "")
}
