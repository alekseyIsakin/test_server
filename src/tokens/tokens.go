package tokens

import (
	"fmt"
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

func ValidRefreshToken(c *gin.Context, refreshToken string, secret []byte) (string, error) {

	token, err := jwt.ParseWithClaims(
		string(refreshToken),
		&RefreshClaims{},
		func(tkn *jwt.Token) (interface{}, error) {
			if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signinig method")
			}
			return secret, nil
		})

	if err != nil {
		return "", err
	}
	now := jwt.TimeFunc().Unix()
	if claim, ok := token.Claims.(*RefreshClaims); ok &&
		token.Valid &&
		claim.ExpiresAt > now {
		return claim.Id, nil
	}

	return "", fmt.Errorf("invalid token")
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

	t, err := token.SignedString([]byte(cfg.GetAccessSecret()))

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

	t, err := r_token.SignedString([]byte(cfg.GetRefreshSecret()))

	return t, err
}
