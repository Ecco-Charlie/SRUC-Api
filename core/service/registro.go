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

type RegistroService struct {
	repository *repository.RegistroRepository
}

func NewRegistroService(db *gorm.DB) *RegistroService {
	rr := repository.NewRegistroRepository(db)
	rr.MigrateDataModels()
	return &RegistroService{
		repository: rr,
	}
}

func (rs *RegistroService) All(params *url.Values) (*[]entity.Registro, *config.Paginator, error) {
	var can int64
	pc := pkg.GetCurrentPage(params)
	query, err := rs.repository.CountAll(params, &can)
	paginator := pkg.GeneratePaginator(int(can), &pc)
	registros, err := rs.repository.All(query, int64(pc))
	if err != nil {
		return nil, nil, err
	}
	return registros, &config.Paginator{
		Cantidad: can,
		Actual:   strconv.Itoa(pc),
		Paginas:  paginator,
	}, nil
}
