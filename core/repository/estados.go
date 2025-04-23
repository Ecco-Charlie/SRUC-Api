package repository

import (
	"net/url"

	"gorm.io/gorm"
	"soft.exe/sruc/core/entity"
)

type EstadosRepository struct {
	db *gorm.DB
}

func NewEstadosRepository(db *gorm.DB) *EstadosRepository {
	return &EstadosRepository{
		db: db,
	}
}

func (er *EstadosRepository) CountAll(params *url.Values, cantidad *int64) (*gorm.DB, error) {
	query := er.db.Model(&entity.Estado{})

	if err := query.Count(cantidad).Error; err != nil {
		return nil, err
	}

	return query, nil
}

func (er *EstadosRepository) All(query *gorm.DB, page int64) (*[]entity.Estado, error) {
	var estados *[]entity.Estado
	offset := int((page - 1) * 11)

	if err := query.Limit(11).Offset(offset).Find(&estados).Error; err != nil {
		return nil, err
	}

	return estados, nil
}

func (er *EstadosRepository) Agregar(estado *entity.Estado) error {
	return er.db.Model(&entity.Estado{}).Create(estado).Error
}

func (er *EstadosRepository) Eliminar(id *int) error {
	return er.db.Delete(&entity.Estado{}, id).Error
}

func (er *EstadosRepository) MigrateDataModel() {}
