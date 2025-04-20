package service

import (
	"encoding/base64"
	"net/url"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"soft.exe/sruc/config"
	"soft.exe/sruc/core/entity"
	"soft.exe/sruc/core/repository"
	"soft.exe/sruc/pkg"
)

type UsuarioService struct {
	repository *repository.UsuarioRespository
}

func NewUsuarioService(db *gorm.DB) *UsuarioService {
	repository := repository.NewUsuarioRepository(db)
	repository.MigrateDataModels()
	return &UsuarioService{
		repository: repository,
	}
}

func (us *UsuarioService) Login(ldto entity.LoginDto) (*entity.Usuario, error) {
	if ldto.NumCuenta == "" || ldto.Password == "" {
		return nil, pkg.ErrBadRequest
	}
	nc, err := strconv.Atoi(ldto.NumCuenta)
	if err != nil {
		return nil, err
	}
	acceso, err := us.repository.FindAccesoByNumCuenta(uint(nc))
	if err != nil {
		return nil, err
	}
	passwd, err := base64.StdEncoding.DecodeString(ldto.Password)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(acceso.Administrativo.Acceso.Password), []byte(passwd)); err != nil {
		return nil, pkg.ErrUnauthorized
	}

	return acceso, nil
}

func (us *UsuarioService) All(params *url.Values) (*[]entity.Usuario, *config.Paginator, error) {
	var can int64
	pc := pkg.GetCurrentPage(params)
	query, err := us.repository.CountAll(params, &can)
	paginator := pkg.GeneratePaginator(int(can), &pc)
	usuarios, err := us.repository.All(query, int64(pc))
	if err != nil {
		return nil, nil, err
	}
	return usuarios, &config.Paginator{
		Cantidad: can,
		Actual:   strconv.Itoa(pc),
		Paginas:  paginator,
	}, nil
}

func (us *UsuarioService) FindByNumCuentaAndRol(NumCuenta *string, rol *string) (*entity.Usuario, error) {
	nc, err := strconv.Atoi(*NumCuenta)
	if err != nil {
		return nil, err
	}
	return us.repository.FindByNumCuentaAndRol(uint(nc), *rol+"s")
}

func (us *UsuarioService) FindExtraData(rol *string, numCuenta string) (*map[string]any, error) {
	nc, err := strconv.Atoi(numCuenta)
	if err != nil {
		return nil, err
	}
	uextra, err := us.repository.FindExtraByNumCuenta((*rol + "s"), uint(nc))
	if err != nil {
		return nil, err
	}
	return uextra, nil
}

func (us *UsuarioService) UpdateUsuario(params *url.Values) error {
	nc, err := strconv.Atoi(params.Get("num_cuenta"))
	if err != nil {
		return err
	}
	am := params.Get("apellmaterno")
	usuario := &entity.Usuario{
		NumCuenta:      uint(nc),
		Nombre:         params.Get("nombre"),
		ApellPaterno:   params.Get("apellpaterno"),
		ApellMaterno:   &am,
		Rol:            params.Get("rol"),
		Administrativo: nil,
		Alumno:         nil,
	}
	switch usuario.Rol {
	case "administrativo":
		usuario.Administrativo = &entity.Administrativo{
			UsuarioNumCuenta: usuario.NumCuenta,
			Area:             params.Get("area"),
		}
	case "alumno":
		usuario.Alumno = &entity.Alumno{
			UsuarioNumCuenta: usuario.NumCuenta,
			Licenciatura:     params.Get("licenciatura"),
		}
	}
	us.repository.EditUsuario(usuario)
	return nil
}
