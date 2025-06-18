package entity

type Registro struct {
	Id     uint
	Inicio CTime
	Fin    CTime

	ProgramaId string
	Programa   Programa

	ServicioId string
	Servicio   Servicio

	ComputadoraId uint
	Computadora   Computadora

	UsuarioId uint
	Usuario   Usuario
}

func (Registro) TableName() string {
	return "registro"
}

type Programa struct {
	Id     string `gorm:"size:20"`
	Nombre string `gorm:"size:25"`
}

func (Programa) TableName() string {
	return "programa"
}

type Servicio struct {
	Id     string `gorm:"size:20"`
	Nombre string `gorm:"size:25"`
}

func (Servicio) TableName() string {
	return "servicio"
}
