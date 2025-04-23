package repository

import (
	"net/url"

	"gorm.io/gorm"
	"soft.exe/sruc/core/entity"
)

type ComputadoraRepository struct {
	db *gorm.DB
}

func NewComputadoraRepository(db *gorm.DB) *ComputadoraRepository {
	return &ComputadoraRepository{
		db: db,
	}
}

func (cr *ComputadoraRepository) CountAll(params *url.Values, cantidad *int64) (*gorm.DB, error) {
	query := cr.db.Model(&entity.Computadora{}).Preload("Estado").Preload("Ubicacion")

	if params.Has("ue") {
		ue := params.Get("ue")
		if ue != "all" {
			query = query.Where("ubicacion_id = ?", ue)
		}
	}

	search := params.Get("search")
	if search != "" {
		query = query.Where(params.Get("type")+" LIKE ?", "%"+search+"%")
	}

	if params.Has("ee") {
		ee := params.Get("ee")
		if ee != "all" {
			query = query.Where("estado_id = ?", ee)
		}
	}

	if err := query.Count(cantidad).Error; err != nil {
		return nil, err
	}

	return query, nil
}

func (cr *ComputadoraRepository) All(query *gorm.DB, page int64) (*[]entity.Computadora, error) {
	var computadoras *[]entity.Computadora
	offset := int((page - 1) * 11)
	if err := query.Limit(11).Offset(offset).Find(&computadoras).Error; err != nil {
		return nil, err
	}
	return computadoras, nil
}

func (cr *ComputadoraRepository) AllUbicaciones() (*[]entity.Ubicacion, error) {
	var ubicaciones *[]entity.Ubicacion
	if err := cr.db.Model(&entity.Ubicacion{}).Select("Id, Nombre").Find(&ubicaciones).Error; err != nil {
		return nil, err
	}
	return ubicaciones, nil
}

func (cr *ComputadoraRepository) AllEstados() (*[]entity.Estado, error) {
	var estados *[]entity.Estado
	if err := cr.db.Model(&entity.Estado{}).Select("Id,Nombre").Find(&estados).Error; err != nil {
		return nil, err
	}
	return estados, nil
}

func (cr *ComputadoraRepository) UpdateEstado(ic, ie *int) error {
	return cr.db.Model(&entity.Computadora{}).Where("id = ?", ic).Update("EstadoId", ie).Error
}

func (cr *ComputadoraRepository) DeleteComputadora(id *int) error {
	return cr.db.Delete(&entity.Computadora{}, id).Error
}

func (cr *ComputadoraRepository) MigrateDataModels() {
	cr.db.AutoMigrate(
		&entity.Ubicacion{},
		&entity.Estado{},
		&entity.Clase{},
		&entity.Computadora{},
	)
}
