package routs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
}
type Authorizer struct {
}

func (a *Authorizer) UserAuthHandler(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &Claims{
		jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Unix() + 60*60*24,
			IssuedAt:  jwt.TimeFunc().Unix(),
		},
	})
	t, err := token.SignedString([]byte("somesecretiugiuuhuygtrdes4567g8h98ojiugytfr6udi7f8giuho;ytrxiy7fuyigubyit6tyuvkey"))
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, t)
}
