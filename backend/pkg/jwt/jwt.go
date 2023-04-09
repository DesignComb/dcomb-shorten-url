package jwt

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"main/config"
	"time"
)

// Key cookie key
const Key = "token"

// Claims Token的結構，裡面放你要的資訊
type Claims struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	jwt.StandardClaims
}

// GenerateToken 產生Token
func GenerateToken(id, email, name, picture string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(config.Val.JWTTokenLife) * time.Second) // Token有效時間

	claims := Claims{
		id,
		email,
		name,
		picture,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "tomato",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(config.Val.JWTSecret))

	return token, err
}

// ParseToken 驗證Token對不對，如果對就回傳user info
func ParseToken(token string) (id, email, name, picture string, err error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Val.JWTSecret), nil
	})

	if err != nil {
		return "", "","", "", err
	}

	claims, ok := tokenClaims.Claims.(*Claims)
	if !ok || !tokenClaims.Valid {
		return "", "","", "", errors.New("tokenClaims invalid")
	}

	return claims.Id, claims.Email, claims.Name, claims.Picture, nil
}
