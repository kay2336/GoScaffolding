package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyCustomClaims struct {
	ID        uint   `json:"id"`
	Username  string `json:"Username"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

// 签名密钥
const signKey = "wadaxinoKay"

// GenerateTokenUsingHS256
// 生成token令牌
func GenerateTokenUsingHS256(userId uint, username string, authority int) (string, error) {
	claim := MyCustomClaims{
		ID:        userId,
		Username:  username,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			//Id:        RandStr(10),                           // wt ID, 类似于盐值
			Issuer:    "Evince_s",                            // 签发者
			Subject:   "User",                                // 签发对象
			Audience:  "user",                                // 签发受众
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // 过期时间
			NotBefore: time.Now().Unix(),                     // 最早使用时间
			IssuedAt:  time.Now().Unix(),                     // 签发时间
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(signKey))
	return token, err
}

// ParseTokenHs256
// 解析token令牌，有奇怪的写法
func ParseTokenHs256(token string) (*MyCustomClaims, error) {
	// 解析token
	tokenClaims, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}

	// 判断token是否有效
	if !tokenClaims.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := tokenClaims.Claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
