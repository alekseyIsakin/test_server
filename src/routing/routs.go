package routs

import (
	"encoding/base64"
	"net/http"
	"test_server/src/config"
	"test_server/src/model"
	"test_server/src/tokens"

	"github.com/gin-gonic/gin"
)

func UserAuthHandler(c *gin.Context) {
	cfg := config.GetConfig()
	uuid := c.Param("uuid")
	atoken, err_a := tokens.GenAccesToken(c, uuid)
	rtoken, err_r := tokens.GenRefreshToken(c, uuid)

	model.UpdateRefreshTokenForUser(c, rtoken, uuid)

	if err_a != nil || err_r != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can`t create token"})
		panic(err_a)
	}

	c.SetCookie(
		"refresh",
		base64.RawURLEncoding.EncodeToString([]byte(rtoken)),
		cfg.GetMaxAgeRefresh(),
		"/",
		cfg.GetDomain(),
		true,
		true)

	c.JSON(http.StatusOK, atoken+cfg.GetTokenDelimiter()+rtoken)
}

func RenewRefreshToken(c *gin.Context) {
	cfg := config.GetConfig()
	cr, err := c.Cookie("refresh")

	if err != nil {
		c.JSON(http.StatusUnauthorized, "wrong refresh token")
		return
	}
	token, _ := base64.RawURLEncoding.DecodeString(cr)

	if err != nil {
		c.JSON(http.StatusUnauthorized, "cant provide new token")
		return
	}

	rtoken, err := model.ReplaceRefreshTokenForUser(c, string(token))
	if err != nil {
		c.JSON(http.StatusUnauthorized, "wrong token"+rtoken)
		return
	}

	c.SetCookie(
		"refresh",
		base64.RawURLEncoding.EncodeToString([]byte(rtoken)),
		cfg.GetMaxAgeRefresh(),
		"/",
		cfg.GetDomain(),
		true,
		true)

	c.JSON(http.StatusOK, "")
}
