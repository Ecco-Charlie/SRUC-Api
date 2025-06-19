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

type ProgramasService struct {
	repository *repository.ProgramasRepository
}

func NewProgramasService(db *gorm.DB) *ProgramasService {
	pr := repository.NewProgramasRepository(db)
	pr.MigrateDataModels()
	return &ProgramasService{
		repository: pr,
	}
}

func (ps *ProgramasService) All(params *url.Values) (*[]entity.Programa, *config.Paginator, error) {
	var can int64
	pc := pkg.GetCurrentPage(params)
	query, err := ps.repository.CountAll(params, &can)
	paginator := pkg.GeneratePaginator(int(can), &pc)
	programas, err := ps.repository.All(query, int64(pc))
	if err != nil {
		return nil, nil, err
	}
	return programas, &config.Paginator{
		Cantidad: can,
		Actual:   strconv.Itoa(pc),
		Paginas:  paginator,
	}, nil
}

func (ps *ProgramasService) AgregarPrograma(params *url.Values) error {
	programa := &entity.Programa{
		Id:     params.Get("paquete"),
		Nombre: params.Get("nombre"),
	}
	return ps.repository.Agregar(programa)
}

func (ps *ProgramasService) ElimnarPrograma(id string) error {
	return ps.repository.Eliminar(&id)
}

func (ps *ProgramasService) AllUnix() (*[]entity.Programa, error) {
	return ps.repository.AllUnix()
}

func (ps *ProgramasService) AllWindows() (*[]entity.Programa, error) {
	return ps.repository.AllWindows()
}
