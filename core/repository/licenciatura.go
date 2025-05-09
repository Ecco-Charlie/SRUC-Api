package repository

import (
	"net/url"

	"gorm.io/gorm"
	"soft.exe/sruc/core/entity"
)

type LicenciaturaRepository struct {
	db *gorm.DB
}

func NewLicenciaturaRepository(db *gorm.DB) *LicenciaturaRepository {
	return &LicenciaturaRepository{
		db: db,
	}
}

func (ls *LicenciaturaRepository) CountAll(params *url.Values, cantidad *int64) (*gorm.DB, error) {
	query := ls.db.Model(&entity.Licenciatura{})

	if err := query.Count(cantidad).Error; err != nil {
		return nil, err
	}

	return query, nil
}

func (ls *LicenciaturaRepository) All(query *gorm.DB, page int64) (*[]entity.Licenciatura, error) {
	var licenciaturas *[]entity.Licenciatura
	offset := int((page - 1) * 11)

	if err := query.Limit(11).Offset(offset).Find(&licenciaturas).Error; err != nil {
		return nil, err
	}

	return licenciaturas, nil
}

func (ls *LicenciaturaRepository) Agregar(licenciatura *entity.Licenciatura) error {
	return ls.db.Create(licenciatura).Error
}

func (ls *LicenciaturaRepository) Eliminar(id *int) error {
	return ls.db.Where("Id = ?", id).Delete(&entity.Licenciatura{}).Error
}

func (lr *LicenciaturaRepository) MigrateDataModels() {}
