package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase(cfg *ApiConfig) *gorm.DB {
	log.Println("Intentando conectarse a la Base de Datos...")
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", *cfg.User, *cfg.Passsword, *cfg.Host, *cfg.Port, *cfg.Database)
	db, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		log.Fatalf("Error al intentar conectarse a la Base de Datos, causa: %s", err)
	}
	log.Println("Base de Datos conectada exitosamente")
	return db
}
