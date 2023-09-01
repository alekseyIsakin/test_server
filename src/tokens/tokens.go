package tokens

import (
	"net/http"
	"test_server/src/config"
	"test_server/src/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
		c.JSON(http.StatusBadRequest, "Wrong user id")
		panic("Wrong user id: " + uuid)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &AccesClaims{
		jwt.StandardClaims{
			Id:        user.UUID,
			IssuedAt:  jwt.TimeFunc().Unix(),
			ExpiresAt: jwt.TimeFunc().Unix() + 60*60*24,
		},
	})

	t, err := token.SignedString([]byte(cfg.GetSecret()))

	return t, err
}

func GenRefreshToken(c *gin.Context, uuid string) (string, error) {
	cfg := config.GetConfig()

	r_token := jwt.NewWithClaims(jwt.SigningMethodHS512, &RefreshClaims{
		jwt.StandardClaims{
			Id:        uuid,
			IssuedAt:  jwt.TimeFunc().Unix(),
			ExpiresAt: jwt.TimeFunc().Unix() + 60*60*24*14,
		},
	})

	t, err := r_token.SignedString([]byte(cfg.GetSecret()))

	return t, err
}