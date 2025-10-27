package config

import "os"

var JWT_SECRET_KEY []byte

func LoadConfig() {
	JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))
}
