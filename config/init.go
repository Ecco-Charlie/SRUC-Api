package config

import (
	"log"

	"soft.exe/sruc/pkg"
)

func LoadApiConfig() *ApiConfig {
	log.Println("Cargando los datos de la aplicación...")
	defer log.Println("Variables cargadas exitosamente")
	return &ApiConfig{
		User:      pkg.GetStrictEnv("DB_USER"),
		Passsword: pkg.GetStrictEnv("DB_PASSWORD"),
		Host:      pkg.GetEnv("MYSQL_HOST", "127.0.0.1"),
		Port:      pkg.GetEnv("MYSQL_PORT", "3306"),
		Database:  pkg.GetEnv("DB_NAME", "sruc"),
		Addr:      pkg.GetEnv("ADDRRESS", ":6767"),
	}
}
