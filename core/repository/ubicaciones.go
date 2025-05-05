package repository

import (
	"net/url"

	"gorm.io/gorm"
	"soft.exe/sruc/core/entity"
)

type UbicacionesRepository struct {
	db *gorm.DB
}

func NewUbicacionesRepository(db *gorm.DB) *UbicacionesRepository {
	return &UbicacionesRepository{
		db: db,
	}
}

func (ur *UbicacionesRepository) CountAll(params *url.Values, cantidad *int64) (*gorm.DB, error) {
	query := ur.db.Model(&entity.Ubicacion{})

	if err := query.Count(cantidad).Error; err != nil {
		return nil, err
	}

	return query, nil
}

func (ur *UbicacionesRepository) All(query *gorm.DB, page int64) (*[]entity.Ubicacion, error) {
	var ubicaciones *[]entity.Ubicacion
	offset := int((page - 1) * 11)

	if err := query.Limit(11).Offset(offset).Find(&ubicaciones).Error; err != nil {
		return nil, err
	}

	return ubicaciones, nil
}

func (ur *UbicacionesRepository) Agregar(ubicacion *entity.Ubicacion) error {
	return ur.db.Create(ubicacion).Error
}

func (ur *UbicacionesRepository) Eliminar(id *int) error {
	return ur.db.Where("Id = ?", id).Delete(&entity.Ubicacion{}).Error
}

func (ur *UbicacionesRepository) AllNature() (*[]entity.Ubicacion, error) {
	var ubicaciones *[]entity.Ubicacion
	if err := ur.db.Find(&ubicaciones).Error; err != nil {
		return nil, err
	}
	return ubicaciones, nil
}

func (ur *UbicacionesRepository) MigrateDataModels() {}
