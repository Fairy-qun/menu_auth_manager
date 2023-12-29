package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"errors"
)

// 定义签名
var signingKey = []byte("hello fairy")


// 定义结构体
type JwtCustClaims struct {
	Id int `json:"id,omitempty`
	Username string `json:"username",omitempty`
	// 继承jwt里面的结构体
	jwt.RegisteredClaims `json:"jwt.RegisteredClaims"`
}

// 生成token
func GenerateToken(id int,username string) (string,error) {
	iJwtCustClaims := 	JwtCustClaims{
		Id: id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// 设置过期时间，这里设置30分钟
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			// 签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 主题(无所谓要不要)
			Subject: "token",
		},
	}
	// 使用
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,iJwtCustClaims)
	return token.SignedString(signingKey)
}

// 解析token
func ParseToken(tokenStr string) (JwtCustClaims, error) {
	iJwtCustClaims := JwtCustClaims{}
	// 解析
	token, err := jwt.ParseWithClaims(tokenStr, &iJwtCustClaims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	// 错误处理
	if err == nil && !token.Valid {
		err = errors.New("Invalid Token")
	}
	return iJwtCustClaims, err
}

// token验证
func TokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	if err != nil {
		return false
	}
	return true
}