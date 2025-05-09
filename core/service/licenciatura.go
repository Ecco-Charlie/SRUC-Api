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

type LicenciaturaService struct {
	repository *repository.LicenciaturaRepository
}

func NewLicenciaturaService(db *gorm.DB) *LicenciaturaService {
	lr := repository.NewLicenciaturaRepository(db)
	lr.MigrateDataModels()
	return &LicenciaturaService{
		repository: lr,
	}
}

func (ls *LicenciaturaService) All(params *url.Values) (*[]entity.Licenciatura, *config.Paginator, error) {
	var can int64
	pc := pkg.GetCurrentPage(params)
	query, err := ls.repository.CountAll(params, &can)
	paginator := pkg.GeneratePaginator(int(can), &pc)
	licenciaturas, err := ls.repository.All(query, int64(pc))
	if err != nil {
		return nil, nil, err
	}
	return licenciaturas, &config.Paginator{
		Cantidad: can,
		Actual:   strconv.Itoa(pc),
		Paginas:  paginator,
	}, nil
}

func (ls *LicenciaturaService) AgregarLicencuatura(params *url.Values) error {
	licenciatura := &entity.Licenciatura{
		Nombre: params.Get("nombre"),
	}
	return ls.repository.Agregar(licenciatura)
}

func (us *LicenciaturaService) EliminarLicenciatura(id string) error {
	iu, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return us.repository.Eliminar(&iu)
}
