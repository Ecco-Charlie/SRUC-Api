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

type ComputadoraService struct {
	repository *repository.ComputadoraRepository
}

func NewComputadoraService(db *gorm.DB) *ComputadoraService {
	cr := repository.NewComputadoraRepository(db)
	cr.MigrateDataModels()
	return &ComputadoraService{
		repository: cr,
	}
}

func (cs *ComputadoraService) All(params *url.Values) (*[]entity.Computadora, *config.Paginator, error) {
	var can int64
	pc := pkg.GetCurrentPage(params)
	query, err := cs.repository.CountAll(params, &can)
	paginator := pkg.GeneratePaginator(int(can), &pc)
	computadoras, err := cs.repository.All(query, int64(pc))
	if err != nil {
		return nil, nil, err
	}
	return computadoras, &config.Paginator{
		Cantidad: can,
		Actual:   strconv.Itoa(pc),
		Paginas:  paginator,
	}, nil
}

func (cs *ComputadoraService) AllUbicaciones() (*[]entity.Ubicacion, error) {
	return cs.repository.AllUbicaciones()
}

func (cs *ComputadoraService) AllEstados() (*[]entity.Estado, error) {
	return cs.repository.AllEstados()
}

func (cs *ComputadoraService) UpdateEstadoComputadora(idComputadora, idEstado string) error {
	ic, err := strconv.Atoi(idComputadora)
	if err != nil {
		return err
	}
	ie, err := strconv.Atoi(idEstado)
	if err != nil {
		return err
	}
	return cs.repository.UpdateEstado(&ic, &ie)
}

func (cs *ComputadoraService) DeleteComputadora(idComputadora string) error {
	id, err := strconv.Atoi(idComputadora)
	if err != nil {
		return err
	}
	return cs.repository.DeleteComputadora(&id)
}

func (cs *ComputadoraService) SaveComputadora(cdto *entity.ComputadoraDto) error {
	c := &entity.Computadora{
		Ip:             cdto.Ip,
		NumPatrimonial: cdto.NumPatrimonial,
		UbicacionId:    uint(cdto.Ubicacion),
		EstadoId:       1,
	}
	return cs.repository.SaveComputadora(c)
}
