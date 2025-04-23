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

type UbicacionesService struct {
	repository *repository.UbicacionesRepository
}

func NewUbicacionesService(db *gorm.DB) *UbicacionesService {
	ur := repository.NewUbicacionesRepository(db)
	ur.MigrateDataModels()
	return &UbicacionesService{
		repository: ur,
	}
}

func (us *UbicacionesService) All(params *url.Values) (*[]entity.Ubicacion, *config.Paginator, error) {
	var can int64
	pc := pkg.GetCurrentPage(params)
	query, err := us.repository.CountAll(params, &can)
	paginator := pkg.GeneratePaginator(int(can), &pc)
	ubicaciones, err := us.repository.All(query, int64(pc))
	if err != nil {
		return nil, nil, err
	}
	return ubicaciones, &config.Paginator{
		Cantidad: can,
		Actual:   strconv.Itoa(pc),
		Paginas:  paginator,
	}, nil
}

func (us *UbicacionesService) AgregarUbicacion(params *url.Values) error {
	cp, err := strconv.Atoi(params.Get("capacidad"))
	if err != nil {
		return err
	}
	ubicacion := &entity.Ubicacion{
		Nombre:      params.Get("nombre"),
		Descripcion: params.Get("descripcion"),
		Capacidad:   uint8(cp),
	}
	return us.repository.Agregar(ubicacion)
}

func (us *UbicacionesService) EliminarUbicacion(id string) error {
	iu, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return us.repository.Eliminar(&iu)
}
