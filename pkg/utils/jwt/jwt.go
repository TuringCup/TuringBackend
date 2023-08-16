package jwt

import (
	"errors"
	"time"

	"github.com/TuringCup/TuringBackend/consts"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	IP       string `json:"ip"`
	jwt.StandardClaims
}

var jwtSecretKey = []byte("njust")

// 生成token
func GenerateToken(id int, username string, ip string) (accessToken, refreshToken string, err error) {
	nowTime := time.Now()
	tokenExpireTime := nowTime.Add(consts.AccessTokenExpireDuration)
	refreshTokenExpireTime := nowTime.Add(consts.RefreshTokenExpireDuration)
	tokenClaims := Claims{
		ID:       id,
		Username: username,
		IP:       ip,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpireTime.Unix(),
			Issuer:    "turingCup",
		},
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims).SignedString(jwtSecretKey)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: refreshTokenExpireTime.Unix(),
		Issuer:    "turingCup",
	}).SignedString(jwtSecretKey)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

// 解析token,解析失败返回nil
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (any, error) {
		return jwtSecretKey, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func VerifyToken(accessToken, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	accessTokenClaims, err := ParseToken(accessToken)
	if err != nil {
		return "", "", errors.New("invalid token")
	}
	refreshTokenClaims, err := ParseToken(refreshToken)
	if err != nil {
		return "", "", errors.New("invalid token")
	}
	if time.Now().Unix() < accessTokenClaims.ExpiresAt {
		return accessToken, refreshToken, nil
	}
	if time.Now().Unix() >= accessTokenClaims.ExpiresAt && time.Now().Unix() <= refreshTokenClaims.ExpiresAt {
		return GenerateToken(accessTokenClaims.ID, accessTokenClaims.Username, accessTokenClaims.IP)
	}
	return "", "", err
}
