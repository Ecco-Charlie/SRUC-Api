package config

import "soft.exe/sruc/pkg"

func LoadApiConfig() *ApiConfig {
	return &ApiConfig{
		User:      pkg.GetStrictEnv("DB_USER"),
		Passsword: pkg.GetStrictEnv("DB_PASSWORD"),
		Host:      pkg.GetEnv("MYSQL_HOST", "127.0.0.1"),
		Port:      pkg.GetEnv("MYSQL_PORT", "3306"),
		Database:  pkg.GetEnv("DB_NAME", "sruc"),
		Addr:      pkg.GetEnv("ADDRRESS", ":6767"),
	}
}
