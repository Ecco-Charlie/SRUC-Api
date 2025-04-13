package repository

import (
	"gorm.io/gorm"
	"soft.exe/sruc/core/entity"
	"soft.exe/sruc/pkg"
)

type UsuarioRespository struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) *UsuarioRespository {
	return &UsuarioRespository{
		db: db,
	}
}

func (ur *UsuarioRespository) FindAccesoByNumCuenta(NumCuenta uint) (*entity.Acceso, error) {
	var acceso *entity.Acceso
	if err := ur.db.Joins("Usuario").First(&acceso, NumCuenta).Error; err != nil {
		return nil, pkg.ErrUserNotFound
	}
	return acceso, nil
}

func (ur *UsuarioRespository) MigrateDataModels() {
	ur.db.AutoMigrate(
		&entity.Usuario{},
		&entity.Administrativo{},
		&entity.Alumno{},
		&entity.Acceso{},
	)
}
