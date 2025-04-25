package repository

import (
	"errors"
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

func (ur *UsuarioRespository) FindAccesoByNumCuenta(NumCuenta uint) (*entity.Usuario, error) {
	var acceso *entity.Usuario
	if err := ur.db.Preload("Administrativo.Acceso").First(&acceso, "num_cuenta = ?", NumCuenta).Error; err != nil {
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
	if err := ur.db.First(&usuario, NumCuenta).Error; err != nil {
		return nil, err
	}
	return usuario, nil
}

func (ur *UsuarioRespository) EditUsuario(usuario *entity.Usuario) error {
	if usuario.Alumno == nil {
		ur.db.Delete(&entity.Alumno{}, usuario.NumCuenta)
	}
	if usuario.Administrativo == nil {
		ur.db.Delete(&entity.Administrativo{}, usuario.NumCuenta)
	} else if usuario.Administrativo.Acceso == nil {
		ur.db.Delete(&entity.Acceso{}, usuario.NumCuenta)
	}

	if err := ur.db.Updates(&usuario).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UsuarioRespository) FindExtraByNumCuenta(tabla *string, NumCuenta uint) (any, error) {
	switch *tabla {
	case "administrativo":
		var administrativo *entity.Administrativo
		if err := ur.db.Model(&entity.Administrativo{}).Preload("Acceso").Where("usuario_num_cuenta = ?", NumCuenta).First(&administrativo).Error; err != nil {
			return nil, err
		}
		return administrativo, nil
	case "alumno":
		var alumno *entity.Alumno
		if err := ur.db.Model(&entity.Alumno{}).Where("usuario_num_cuenta = ?", NumCuenta).First(&alumno).Error; err != nil {
			return nil, err
		}
		return alumno, nil
	default:
		return nil, errors.New("empty")
	}
}

func (ur *UsuarioRespository) DeleteUsuarioByNumCuenta(NumCuenta uint) error {
	if err := ur.db.Delete(&entity.Usuario{}, NumCuenta).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UsuarioRespository) MigrateDataModels() {
	ur.db.AutoMigrate(
		&entity.Usuario{},
		&entity.Administrativo{},
		&entity.Alumno{},
		&entity.Acceso{},
	)
}
