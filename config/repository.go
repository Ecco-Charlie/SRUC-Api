package config

type Repository interface {
	MigrateDataModels()
}
