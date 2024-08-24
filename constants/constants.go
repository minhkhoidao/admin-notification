package constants

import "os"

var (
	Secret             = os.Getenv("SECRET")
	AccessTokenExpiry  = os.Getenv("ACCESS_TOKEN_EXPIRE")
	RefreshTokenExpiry = os.Getenv("REFRESH_TOKEN_EXPIRE")
)
