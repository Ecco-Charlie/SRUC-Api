package repository

import (
	"net/url"

	"gorm.io/gorm"
	"soft.exe/sruc/core/entity"
)

type ProgramasRepository struct {
	db *gorm.DB
}

func NewProgramasRepository(db *gorm.DB) *ProgramasRepository {
	return &ProgramasRepository{
		db: db,
	}
}

func (pr *ProgramasRepository) CountAll(params *url.Values, cantidad *int64) (*gorm.DB, error) {
	query := pr.db.Model(&entity.Programa{})

	if err := query.Count(cantidad).Error; err != nil {
		return nil, err
	}

	return query, nil
}

func (pr *ProgramasRepository) All(query *gorm.DB, page int64) (*[]entity.Programa, error) {
	var estados *[]entity.Programa
	offset := int((page - 1) * 11)

	if err := query.Limit(11).Offset(offset).Find(&estados).Error; err != nil {
		return nil, err
	}

	return estados, nil
}

func (pr *ProgramasRepository) Agregar(programa *entity.Programa) error {
	return pr.db.Create(programa).Error
}

func (pr *ProgramasRepository) Eliminar(id *string) error {
	return pr.db.Delete(&entity.Programa{}, id).Error
}

func (pr *ProgramasRepository) MigrateDataModels() {}
