package repository

import (
	"net/url"

	"gorm.io/gorm"
	"soft.exe/sruc/core/entity"
)

type AreaRepository struct {
	db *gorm.DB
}

func NewAreaRepository(db *gorm.DB) *AreaRepository {
	return &AreaRepository{
		db: db,
	}
}

func (ar *AreaRepository) CountAll(params *url.Values, cantidad *int64) (*gorm.DB, error) {
	query := ar.db.Model(&entity.Area{})

	if err := query.Count(cantidad).Error; err != nil {
		return nil, err
	}

	return query, nil
}

func (ar *AreaRepository) All(query *gorm.DB, page int64) (*[]entity.Area, error) {
	var areas *[]entity.Area
	offset := int((page - 1) * 11)

	if err := query.Limit(11).Offset(offset).Find(&areas).Error; err != nil {
		return nil, err
	}

	return areas, nil
}

func (ar *AreaRepository) Agregar(area *entity.Area) error {
	return ar.db.Create(area).Error
}

func (ar *AreaRepository) Eliminar(id *int) error {
	return ar.db.Where("Id = ?", id).Delete(&entity.Area{}).Error
}

func (ar *AreaRepository) MigrateDataModels() {}
