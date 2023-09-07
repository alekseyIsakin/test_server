package tokens

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	g_uuid "github.com/google/uuid"
	"net/http"
	"test_server/src/config"
	"test_server/src/model"
)

type AccesClaims struct {
	jwt.StandardClaims
}

type RefreshClaims struct {
	jwt.StandardClaims
}

func GenAccesToken(c *gin.Context, uuid string) (string, error) {
	cfg := config.GetConfig()

	user := model.FindUserByUUID(c, uuid)

	if user.UUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong user id"})
		return "", fmt.Errorf("wrong user id: %s", uuid)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &AccesClaims{
		jwt.StandardClaims{
			Id:        user.UUID,
			IssuedAt:  jwt.TimeFunc().Unix(),
			ExpiresAt: jwt.TimeFunc().Unix() + 60*60*24,
		},
	})

	t, err := token.SignedString([]byte(cfg.GetAccessSecret()))

	return t, err
}

func GenRefreshToken(c *gin.Context, uuid string) (string, error) {
	cfg := config.GetConfig()
	x := g_uuid.New()

	return uuid + cfg.GetTokenDelimiter() + x.String(), nil
}
