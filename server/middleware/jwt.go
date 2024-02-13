package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sxz799/surveyX/model/common/response"
	"time"
)

var JwtKey = []byte("survey_secret_key")

type UserClaims struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	jwt.RegisteredClaims
}

func GenToken(userId int, username, nickname string) (tokenStr string, err error) {
	// 创建jwt accessClaims 设置过期时间25s
	claims := &UserClaims{
		UserId:   userId,
		Username: username,
		Nickname: nickname,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
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

func ParseToken(c *gin.Context) (claims *UserClaims, err error) {
	tokenStr, err := c.Cookie("token")
	claims = &UserClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := ParseToken(c)
		if err != nil {
			response.FailWithMessage("Token Expired!", c)
			c.Abort()
			return
		}
		if claims != nil {
			if claims.ExpiresAt.Unix()-time.Now().Unix() < 15 {
				str, _ := GenToken(claims.UserId, claims.Username, claims.Nickname)
				c.SetCookie("token", str, 60*30, "", "", false, true)
			}
			c.Set("claims", claims)
			c.Next()
		} else {
			response.FailWithMessage("Token Expired", c)
			c.Abort()
			return
		}

	}

}
