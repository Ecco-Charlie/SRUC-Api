package repository

import (
	"net/url"

	"gorm.io/gorm"
	"soft.exe/sruc/core/entity"
)

type RegistroRepository struct {
	db *gorm.DB
}

func NewRegistroRepository(db *gorm.DB) *RegistroRepository {
	return &RegistroRepository{
		db: db,
	}
}

func (rr *RegistroRepository) CountAll(params *url.Values, cantidad *int64) (*gorm.DB, error) {
	query := rr.db.Model(&entity.Registro{}).Preload("Usuario").Preload("Computadora").Preload("Computadora.Ubicacion").Preload("Programa").Preload("Servicio")

	if params.Has("inicio") || params.Has("finc") {
		inicio := params.Get("inicio")
		fin := params.Get("fin")

		if inicio != "" && fin != "" {
			query = query.Where("inicio BETWEEN ? AND ?", inicio, fin)
		} else if inicio != "" {
			query = query.Where("inicio > ?", inicio)
		} else if fin != "" {
			query = query.Where("fin < ?", fin)
		}
	}

	if params.Has("hinicio") || params.Has("hfin") {
		hinicio := params.Get("hinicio")
		hfin := params.Get("hfin")

		if hinicio != "" && hfin != "" {
			query = query.Where("TIME(inicio) BETWEEN ? AND ?", hinicio, hfin)
		} else if hinicio != "" {
			query = query.Where("TIME(inicio) > ?", hinicio)
		} else if hfin != "" {
			query = query.Where("TIME(fin) < ?", hfin)
		}
	}

	if params.Has("mes") && params.Get("mes") != "all" {
		query = query.Where("MONTH(inicio) = ?", params.Get("mes"))
	}

	if params.Has("dia") && params.Get("dia") != "all" {
		query = query.Where("DAYOFWEEK(inicio) = ?", params.Get("dia"))
	}

	if params.Has("licenciatura") && params.Get("licenciatura") != "all" {
		query = query.Joins("Usuario").
			Joins("Usuario.Alumno").
			Where("licenciatura = ?", params.Get("licenciatura")).
			Preload("Usuario.Alumno")
	}

	if params.Has("alumno") {
		query = query.Joins("Usuario").
			Where("rol = ?", "alumno")
	}

	if params.Has("administrativo") {
		query = query.Joins("Usuario").
			Where("rol = ?", "administrativo")
	}

	if params.Has("area") && params.Get("area") != "all" {
		query = query.Joins("Usuario").
			Joins("Usuario.Administrativo").
			Where("area = ?", params.Get("area"))
	}

	if err := query.Count(cantidad).Error; err != nil {
		return nil, err
	}

	return query, nil
}

func (rr *RegistroRepository) All(query *gorm.DB, page int64) (*[]entity.Registro, error) {
	var computadoras *[]entity.Registro
	offset := int((page - 1) * 11)

	if err := query.Limit(11).Offset(offset).Find(&computadoras).Error; err != nil {
		return nil, err
	}

	return computadoras, nil
}

func (rr *RegistroRepository) AllLicenciaturas() (*[]string, error) {
	var licenciaturas *[]string
	if err := rr.db.Model(&entity.Alumno{}).Distinct("Licenciatura").Find(&licenciaturas).Error; err != nil {
		return nil, err
	}
	return licenciaturas, nil
}

func (rr *RegistroRepository) AllAreas() (*[]string, error) {
	var areas *[]string
	if err := rr.db.Model(&entity.Administrativo{}).Distinct("Area").Find(&areas).Error; err != nil {
		return nil, err
	}
	return areas, nil
}

func (rr *RegistroRepository) MigrateDataModels() {
	rr.db.AutoMigrate(&entity.Registro{})
}
