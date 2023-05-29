package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyCustomClaims struct {
	UserID   uint
	Username string
	jwt.RegisteredClaims
}

// 签名密钥
const signKey = "wadaxinoKay"

// GenerateTokenUsingHS256
// 生成token令牌
func GenerateTokenUsingHS256(userId uint, username string) (string, error) {
	claim := MyCustomClaims{
		UserID:   userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Evince_s",                                      // 签发者
			Subject:   "User",                                          // 签发对象
			Audience:  jwt.ClaimStrings{"Android_APP", "IOS_APP"},      // 签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),   // 过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), // 最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                  // 签发时间
			ID:        RandStr(10),                                     // wt ID, 类似于盐值
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(signKey))
	return token, err
}

// ParseTokenHs256
// 解析token令牌
func ParseTokenHs256(token string) (*MyCustomClaims, error) {
	// 奇怪的写法
	tokenClaims, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}

	// 判断tokenClaims是否有效
	if !tokenClaims.Valid {
		return nil, errors.New("claim invalid")
	}

	// 没懂
	claims, ok := tokenClaims.Claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
