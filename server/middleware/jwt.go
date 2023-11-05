package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sxz799/surveyX/model/common/response"
	"time"
)

var JwtKey = []byte("survey_secret_key")

type Claims struct {
	Key string `json:"key"`
	jwt.RegisteredClaims
}

func GenToken(key string) (tokenStr string, err error) {
	// 创建jwt accessClaims 设置过期时间25s
	claims := &Claims{
		Key: key,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(JwtKey)
	if err != nil {
		// 创建Token失败
		return "", err
	}
	return tokenStr, nil
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("token")
		if err != nil {
			// 获取access-token失败
			response.FailWithMessage("未获取到登录信息!", c)
			c.Abort()
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
		if err != nil {
			response.FailWithMessage("Token不正确!", c)
			c.Abort()
			return
		}
		if token.Valid {
			if claims.ExpiresAt.Unix()-time.Now().Unix() < 15 {
				str, _ := GenToken(claims.Key)
				c.SetCookie("token", str, 3600, "", "", false, true)
			}
			c.Set("claims", claims)
			c.Next()
		} else {
			response.FailWithMessage("Token已过期，请重新登录！", c)
			c.Abort()
			return
		}

	}

}
