package pkg

import (
	"log"
	"os"
)

func GetEnv(key, def string) *string {
	if value, ok := os.LookupEnv(key); ok {
		return &value
	}
	return &def
}

func GetStrictEnv(key string) *string {
	if value, ok := os.LookupEnv(key); ok {
		return &value
	}
	log.Fatalf("Por favor especificar la variable de entorno %s", key)
	return nil
}
