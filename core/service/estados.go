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

type EstadosService struct {
	repository *repository.EstadosRepository
}

func NewEstadoService(db *gorm.DB) *EstadosService {
	er := repository.NewEstadosRepository(db)
	er.MigrateDataModel()
	return &EstadosService{
		repository: er,
	}
}

func (es *EstadosService) All(params *url.Values) (*[]entity.Estado, *config.Paginator, error) {
	var can int64
	pc := pkg.GetCurrentPage(params)
	query, err := es.repository.CountAll(params, &can)
	paginator := pkg.GeneratePaginator(int(can), &pc)
	estados, err := es.repository.All(query, int64(pc))
	if err != nil {
		return nil, nil, err
	}
	return estados, &config.Paginator{
		Cantidad: can,
		Actual:   strconv.Itoa(pc),
		Paginas:  paginator,
	}, nil
}

func (es *EstadosService) AgregarEstado(params *url.Values) error {
	dpn, err := strconv.Atoi(params.Get("disponibilidad"))
	if err != nil {
		return err
	}
	estado := &entity.Estado{
		Nombre:         params.Get("nombre"),
		Disponibilidad: uint8(dpn),
	}
	return es.repository.Agregar(estado)
}

func (es *EstadosService) EliminarEstado(id string) error {
	ie, err := strconv.Atoi(id)
	if err != nil {
		return nil
	}
	return es.repository.Eliminar(&ie)
}
