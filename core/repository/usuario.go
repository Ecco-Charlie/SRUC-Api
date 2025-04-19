package repository

import (
	"net/url"

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

func (ur *UsuarioRespository) CountAll(params *url.Values, cantidad *int64) (*gorm.DB, error) {
	query := ur.db.Model(&entity.Usuario{})

	if params.Has("er") {
		er := params.Get("er")
		if er != "all" {
			query = query.Where("rol = ?", er)
		}
	}

	if params.Get("search") != "" {
		query = query.Where(params.Get("type")+" LIKE ?", "%"+params.Get("search")+"%")
	}

	if err := query.Count(cantidad).Error; err != nil {
		return nil, err
	}

	return query, nil
}

func (ur *UsuarioRespository) All(query *gorm.DB, page int64) (*[]entity.Usuario, error) {
	var usuarios *[]entity.Usuario
	offset := int((page - 1) * 11)
	if err := query.Limit(11).Offset(offset).Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (ur *UsuarioRespository) FindByNumCuentaAndRol(NumCuenta uint, rol string) (*entity.Usuario, error) {
	var usuario *entity.Usuario
	if err := ur.db.Preload("Administrativo").Preload("Alumno").First(&usuario, NumCuenta).Error; err != nil {
		return nil, err
	}
	return usuario, nil
}

func (ur *UsuarioRespository) MigrateDataModels() {
	ur.db.AutoMigrate(
		&entity.Usuario{},
		&entity.Administrativo{},
		&entity.Alumno{},
		&entity.Acceso{},
	)
}
