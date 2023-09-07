package routs

import (
	"encoding/base64"
	"net/http"
	"strings"
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

	if err_a != nil || err_r != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can`t create token"})
		panic(err_a)
	}

	if ok := model.UpdateRefreshTokenForUser(c, rtoken, uuid); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can`t update token"})
	}

	c.SetCookie(
		"refresh",
		base64.RawURLEncoding.EncodeToString([]byte(rtoken)),
		cfg.GetMaxAgeRefresh(),
		"/",
		cfg.GetDomain(),
		true,
		true)

	c.JSON(http.StatusOK, gin.H{"success": atoken})
}

func RenewRefreshToken(c *gin.Context) {
	cfg := config.GetConfig()
	cr, err := c.Cookie("refresh")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token are not provided"})
		return
	}

	token, _ := base64.RawURLEncoding.DecodeString(cr)

	if err := model.ValidateRToken(c, string(token)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token is not valid"})
		return
	}

	uuid := strings.Split(string(token), cfg.GetTokenDelimiter())[0]

	atoken, err := tokens.GenAccesToken(c, uuid)
	rtoken, err := tokens.GenRefreshToken(c, uuid)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "cant create token"})
		return
	}
	model.UpdateRefreshTokenForUser(c, rtoken, uuid)

	c.SetCookie(
		"refresh",
		base64.RawURLEncoding.EncodeToString([]byte(rtoken)),
		cfg.GetMaxAgeRefresh(),
		"/",
		cfg.GetDomain(),
		true,
		true)

	c.JSON(http.StatusOK, gin.H{"success": atoken})
}
