package consts

import "time"

const (
	AccessTokenExpireDuration  = 5 * time.Minute
	RefreshTokenExpireDuration = 2 * time.Hour
)
