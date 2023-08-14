package jwt

import (
	"github.com/TuringCup/TuringBackend/consts"
	"github.com/dgrijalva/jwt-go"
	"time"
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
	return accessToken, refreshToken, err
}

// 解析token
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
		return "", "", err
	}
	refreshTokenClaims, err := ParseToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	if accessTokenClaims.ExpiresAt > time.Now().Unix() {
		return GenerateToken(accessTokenClaims.ID, accessTokenClaims.Username, accessTokenClaims.IP)
	}
	if refreshTokenClaims.ExpiresAt > time.Now().Unix() {
		return GenerateToken(accessTokenClaims.ID, accessTokenClaims.Username, accessTokenClaims.IP)
	}
	return "", "", err
}
