package service

import (
	"net/url"
	"strconv"

	"gorm.io/gorm"
	"soft.exe/sruc/config"
	"soft.exe/sruc/core/entity"
	"soft.exe/sruc/core/repository"
	"soft.exe/sruc/pkg"
)

type AreaService struct {
	repository *repository.AreaRepository
}

func NewAreaService(db *gorm.DB) *AreaService {
	ar := repository.NewAreaRepository(db)
	ar.MigrateDataModels()
	return &AreaService{
		repository: ar,
	}
}

func (ar *AreaService) All(params *url.Values) (*[]entity.Area, *config.Paginator, error) {
	var can int64
	pc := pkg.GetCurrentPage(params)
	query, err := ar.repository.CountAll(params, &can)
	paginator := pkg.GeneratePaginator(int(can), &pc)
	areas, err := ar.repository.All(query, int64(pc))
	if err != nil {
		return nil, nil, err
	}
	return areas, &config.Paginator{
		Cantidad: can,
		Actual:   strconv.Itoa(pc),
		Paginas:  paginator,
	}, nil
}

func (ar *AreaService) AgregarArea(params *url.Values) error {
	area := &entity.Area{
		Nombre: params.Get("nombre"),
	}
	return ar.repository.Agregar(area)
}

func (ar *AreaService) EliminarArea(id string) error {
	iu, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return ar.repository.Eliminar(&iu)
}
