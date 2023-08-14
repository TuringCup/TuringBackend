package jwt

import (
	"testing"
)

func TestGenerateAndParseToken(t *testing.T) {
	id := 3
	username := "testJwt"
	ip := "127.0.0.1"
	accessToken, _, err := GenerateToken(id, username, ip)
	if err != nil {
		return
	}
	accessTokenClaims, err := ParseToken(accessToken)
	if accessTokenClaims != nil {
		if id != accessTokenClaims.ID {
			t.Errorf("expect %d,but get %d", id, accessTokenClaims.ID)
		}
		if username != accessTokenClaims.Username {
			t.Errorf("expect %s,but get %s", username, accessTokenClaims.Username)
		}
		if ip != accessTokenClaims.IP {
			t.Errorf("expect %s,but get %s", ip, accessTokenClaims.IP)
		}
	} else {
		t.Errorf("parse failed")
	}
	jwtSecretKey = []byte("njustwrong")
	accessTokenClaims, err = ParseToken(accessToken)
	if accessTokenClaims != nil {
		if id != accessTokenClaims.ID {
			t.Errorf("expect %d,but get %d", id, accessTokenClaims.ID)
		}
		if username != accessTokenClaims.Username {
			t.Errorf("expect %s,but get %s", username, accessTokenClaims.Username)
		}
		if ip != accessTokenClaims.IP {
			t.Errorf("expect %s,but get %s", ip, accessTokenClaims.IP)
		}
	}
}
