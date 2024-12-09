package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"strings"
	"time"
)

var JwtKey = []byte("survey_secret_key")
var JwtExpire = time.Hour * 2

type Claims struct {
	UserInfo entity.User `json:"user_info"`
	jwt.RegisteredClaims
}

func GenToken(userInfo entity.User, rememberMe bool) (tokenStr string, err error) {
	var tJwtExpire time.Duration
	if rememberMe {
		tJwtExpire = time.Hour * 24 * 7
	} else {
		tJwtExpire = JwtExpire / 2
	}
	// 创建jwt accessClaims 设置过期时间25s
	userInfo.Password = ""
	claim := &Claims{
		UserInfo: userInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tJwtExpire * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err = token.SignedString(JwtKey)
	if err != nil {
		// 创建Token失败
		return "", err
	}
	return tokenStr, nil
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		tokenStr := strings.Replace(auth, "Bearer ", "", 1)
		var claims = &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
		if err != nil {
			response.FailWithExpire(c)
			c.Abort()
			return
		}
		if token.Valid {
			c.Set("userInfo", claims.UserInfo)
			c.Next()
		} else {
			response.FailWithExpire(c)
		}
	}

}
