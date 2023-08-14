package consts

import "time"

const (
	AccessTokenExpireDuration  = 2 * time.Hour
	RefreshTokenExpireDuration = 10 * 24 * time.Hour
)
