package conf

import (
	"time"
)

type JWTConfig struct {
	Expires time.Duration
	Issuer  string
	Secret  string
}

func GetJWTConfig() JWTConfig {
	return jwtConfig
}
